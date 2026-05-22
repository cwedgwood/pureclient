// Package auth provides authentication helpers for the FlashArray REST API.
//
// FlashArray's primary auth model is a two-token dance:
//
//   - A long-lived API token (issued per-user by an array admin) is sent
//     once to POST /api/2.26/login to obtain a session token. The
//     /api/2.26/ path is bound to the version of the generated client
//     this package calls into (PostApi226LoginWithResponse); if the
//     client is regenerated against a different API version, this
//     package and that login call need to be updated together.
//   - A short-lived session token returned in the X-Auth-Token response
//     header is sent in the x-auth-token request header on every
//     subsequent call. The array invalidates it after a period of
//     inactivity, at which point requests start returning 401.
//
// SessionTokenSource holds the API token and caches the session token,
// refreshing it on demand. Transport wraps an http.RoundTripper to
// inject the session token on every request and transparently re-login
// + retry once on a 401.
//
// The shape mirrors golang.org/x/oauth2's TokenSource + Transport but
// implements FlashArray's own session-token protocol - there is no
// dependency on golang.org/x/oauth2 and the on-the-wire details are
// entirely different.
//
// For FlashArray's separate OAuth2 client-credentials flow (POST
// /oauth2/1.0/token with an RFC 7523 JWT bearer assertion), this
// package is not the right fit; use a custom http.RoundTripper that
// sets Authorization: Bearer. See MIGRATION.md.
package auth

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	pureclient "github.com/cwedgwood/pureclient/client"
)

// SessionTokenSource holds a long-lived FlashArray API token and a
// cached short-lived x-auth-token session token. The zero value is not
// useful: Endpoint and APIToken must be set before use.
//
// Safe for concurrent use.
type SessionTokenSource struct {
	// Endpoint is the array base URL (no trailing slash), e.g.
	// "https://array-vip".
	Endpoint string

	// APIToken is the long-lived FlashArray API token issued by an
	// array admin. Never sent on data requests; only sent in the
	// api-token header to POST /api/X.Y/login during refresh.
	APIToken string

	// HTTPClient is used to perform the login call. It MUST NOT
	// itself inject the session token (avoid recursion); typically
	// it wraps the same *http.Transport you would otherwise use,
	// without any auth wrapping.
	//
	// If nil, http.DefaultClient is used.
	HTTPClient *http.Client

	mu      sync.Mutex
	session string
}

// Compile-time assertion that *SessionTokenSource satisfies TokenSource.
var _ TokenSource = (*SessionTokenSource)(nil)

// maxArrayErrorBodyBytes caps how many bytes of the array's 401
// response body we read when surfacing it in a refresh-failure
// error. The array returns short JSON errors here, so this is well
// above the typical payload; the cap exists only to keep a
// pathological response from blowing up an error string.
const maxArrayErrorBodyBytes = 512

// Token returns the cached session token, calling Refresh if the cache
// is empty.
func (s *SessionTokenSource) Token(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.session == "" {
		if err := s.refreshLocked(ctx); err != nil {
			return "", err
		}
	}
	return s.session, nil
}

// Refresh forces a re-login and replaces the cached session token.
// Transport calls RefreshIfStale on a 401; callers can call Refresh
// directly to pre-warm the cache or to recover from out-of-band
// invalidation.
func (s *SessionTokenSource) Refresh(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.refreshLocked(ctx)
}

// RefreshIfStale logs in only if the cached session token still
// matches stale. Returns the current session token, which may be the
// result of a refresh by another goroutine that won the race. Intended
// for callers that observed a 401 with a particular cached token and
// want to avoid stampeding the login endpoint when many goroutines
// hit the same expired token at once.
func (s *SessionTokenSource) RefreshIfStale(ctx context.Context, stale string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.session != stale {
		return s.session, nil
	}
	if err := s.refreshLocked(ctx); err != nil {
		return "", err
	}
	return s.session, nil
}

func (s *SessionTokenSource) refreshLocked(ctx context.Context) error {
	httpClient := s.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c, err := pureclient.NewClientWithResponses(s.Endpoint,
		pureclient.WithHTTPClient(httpClient),
	)
	if err != nil {
		return err
	}
	resp, err := c.PostApi226LoginWithResponse(ctx, &pureclient.PostApi226LoginParams{
		ApiToken: &s.APIToken,
	})
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		body := resp.Body
		if len(body) > maxArrayErrorBodyBytes {
			body = body[:maxArrayErrorBodyBytes]
		}
		return fmt.Errorf("login failed: %s: %s", resp.Status(), string(body))
	}
	token := resp.HTTPResponse.Header.Get("X-Auth-Token")
	if token == "" {
		return errors.New("login succeeded but X-Auth-Token header was empty")
	}
	s.session = token
	return nil
}

// TokenSource is the minimal interface Transport needs to obtain and
// refresh a session token. *SessionTokenSource satisfies it; callers
// with bespoke token-acquisition requirements (e.g. fetching the API
// token from a secret store on each refresh, or fronting a different
// auth backend that returns a FlashArray-compatible session token)
// can supply their own implementation.
//
// Implementations must be safe for concurrent use.
type TokenSource interface {
	// Token returns the current session token, performing an initial
	// login if none is cached.
	Token(ctx context.Context) (string, error)

	// RefreshIfStale returns a fresh session token, but only performs
	// a re-login if the currently-cached token still equals stale.
	// This lets Transport coalesce concurrent 401s into a single
	// re-login: the first caller refreshes, later callers observe
	// that the cached token has already moved past their stale value
	// and reuse it.
	RefreshIfStale(ctx context.Context, stale string) (string, error)
}

// Transport is an http.RoundTripper that injects the cached session
// token on every request and transparently re-logs-in + retries once
// on a 401.
//
// Requests whose Body is non-replayable (Body != nil and GetBody nil)
// are NOT retried on a 401 - the 401 is surfaced to the caller. This
// matches the behaviour expected of any retry-on-failure transport:
// retrying with an empty/consumed body would produce a worse outcome
// than the original error.
type Transport struct {
	// Base is the underlying http.RoundTripper used to perform the
	// actual HTTP call. If nil, http.DefaultTransport is used.
	Base http.RoundTripper

	// Source supplies and refreshes the session token. Must be
	// non-nil. *SessionTokenSource is the standard implementation;
	// any TokenSource will do.
	Source TokenSource
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Source == nil {
		return nil, errors.New("auth.Transport: Source is nil")
	}
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}

	token, err := t.Source.Token(req.Context())
	if err != nil {
		return nil, err
	}
	first := req.Clone(req.Context())
	first.Header.Set("x-auth-token", token)
	resp, err := base.RoundTrip(first)
	if err != nil || resp.StatusCode != http.StatusUnauthorized {
		return resp, err
	}

	// Token was rejected. If we cannot replay the body, surface the
	// 401 unmodified - better than retrying with an empty body.
	if req.Body != nil && req.GetBody == nil {
		return resp, nil
	}

	// Close the 401 before the retry. The body is typically short
	// (a JSON error from the array); not draining a few hundred
	// bytes is fine - Close releases the connection back to the
	// pool just as well for short responses.
	//
	// Exception: if the refresh below fails, the original 401 body
	// is the most useful piece of diagnostic information the array
	// gave us - losing it would leave the caller with a generic
	// "re-login failed" wrapping a generic refresh error. So
	// peek-then-close instead of close-then-refresh.
	origBody, _ := io.ReadAll(io.LimitReader(resp.Body, maxArrayErrorBodyBytes))
	resp.Body.Close()

	// Re-login only if no other goroutine has already rotated past
	// the token we just saw rejected. RefreshIfStale collapses N
	// concurrent 401s on a shared stale token into one login.
	newToken, err := t.Source.RefreshIfStale(req.Context(), token)
	if err != nil {
		if trimmed := bytes.TrimSpace(origBody); len(trimmed) > 0 {
			return nil, fmt.Errorf("re-login after 401 (original 401 body: %q): %w", string(trimmed), err)
		}
		return nil, fmt.Errorf("re-login after 401: %w", err)
	}
	retry := req.Clone(req.Context())
	if req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			return nil, fmt.Errorf("auth.Transport: GetBody for 401 retry: %w", err)
		}
		retry.Body = body
	}
	retry.Header.Set("x-auth-token", newToken)
	return base.RoundTrip(retry)
}
