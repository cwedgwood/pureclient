package auth

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// fakeArray emulates the FlashArray endpoints that auth touches:
//   - POST /api/2.26/login: returns 200 + X-Auth-Token=<sessionTok>
//     when api-token header matches expectedAPIToken; else 401.
//   - everything else: returns 200 if x-auth-token matches the current
//     session token, else 401.
//
// loginCalls / dataCalls track how often each path was hit.
type fakeArray struct {
	expectedAPIToken string
	sessionTok       atomic.Value // string
	loginCalls       atomic.Int32
	dataCalls        atomic.Int32

	// If true, each successful login generates a new session token
	// (sess-rotated-N) instead of returning the cached sessionTok.
	// Models a FlashArray that invalidates the prior session on
	// every new login - the strict reading of session semantics.
	rotateOnLogin bool
	rotateCounter atomic.Int32

	// dataBodies captures the request body of every data call (in
	// the order they arrived). Used by body-replay tests.
	bodiesMu   sync.Mutex
	dataBodies [][]byte

	// data401Body, if non-empty, is written as the response body
	// alongside a 401 status on data-path rejections. Used to
	// verify that callers can recover the array's original 401
	// diagnostic message when refresh subsequently fails.
	data401Body []byte

	server *httptest.Server
}

func newFakeArray(t *testing.T, apiTok, initialSession string) *fakeArray {
	t.Helper()
	f := &fakeArray{expectedAPIToken: apiTok}
	f.sessionTok.Store(initialSession)
	f.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.26/login" && r.Method == http.MethodPost {
			f.loginCalls.Add(1)
			if r.Header.Get("api-token") != f.expectedAPIToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if f.rotateOnLogin {
				next := fmt.Sprintf("sess-rotated-%d", f.rotateCounter.Add(1))
				f.sessionTok.Store(next)
			}
			w.Header().Set("X-Auth-Token", f.sessionTok.Load().(string))
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"items":[]}`))
			return
		}
		f.dataCalls.Add(1)
		body, _ := io.ReadAll(r.Body)
		f.bodiesMu.Lock()
		f.dataBodies = append(f.dataBodies, body)
		f.bodiesMu.Unlock()
		if r.Header.Get("x-auth-token") != f.sessionTok.Load().(string) {
			w.WriteHeader(http.StatusUnauthorized)
			if len(f.data401Body) > 0 {
				_, _ = w.Write(f.data401Body)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"items":[]}`))
	}))
	t.Cleanup(f.server.Close)
	return f
}

func (f *fakeArray) rotateSession(s string) { f.sessionTok.Store(s) }

func TestSessionTokenSource_TokenCachesAfterFirstRefresh(t *testing.T) {
	f := newFakeArray(t, "api-tok-good", "sess-1")
	s := &SessionTokenSource{
		Endpoint:   f.server.URL,
		APIToken:   "api-tok-good",
		HTTPClient: f.server.Client(),
	}
	tok1, err := s.Token(context.Background())
	if err != nil {
		t.Fatalf("first Token: %v", err)
	}
	tok2, err := s.Token(context.Background())
	if err != nil {
		t.Fatalf("second Token: %v", err)
	}
	if tok1 != "sess-1" || tok2 != "sess-1" {
		t.Errorf("got tokens %q, %q; want both sess-1", tok1, tok2)
	}
	if got := f.loginCalls.Load(); got != 1 {
		t.Errorf("login called %d times; want 1 (cache should hold)", got)
	}
}

func TestSessionTokenSource_BadAPITokenFails(t *testing.T) {
	f := newFakeArray(t, "api-tok-good", "sess-1")
	s := &SessionTokenSource{
		Endpoint:   f.server.URL,
		APIToken:   "api-tok-bad",
		HTTPClient: f.server.Client(),
	}
	_, err := s.Token(context.Background())
	if err == nil {
		t.Fatal("expected error from bad API token; got nil")
	}
	if !strings.Contains(err.Error(), "login failed") {
		t.Errorf("error text %q does not mention 'login failed'", err.Error())
	}
}

func TestTransport_InjectsTokenAndRetriesOn401(t *testing.T) {
	f := newFakeArray(t, "api-tok-good", "sess-1")
	s := &SessionTokenSource{
		Endpoint:   f.server.URL,
		APIToken:   "api-tok-good",
		HTTPClient: f.server.Client(),
	}
	// Pre-warm the source with a known-good token so the first data
	// call succeeds without a login.
	if _, err := s.Token(context.Background()); err != nil {
		t.Fatalf("pre-warm: %v", err)
	}
	tr := &Transport{Base: f.server.Client().Transport, Source: s}
	httpClient := &http.Client{Transport: tr}

	// First call: should succeed using the cached token.
	resp1, err := httpClient.Get(f.server.URL + "/api/2.26/volumes")
	if err != nil {
		t.Fatalf("first GET: %v", err)
	}
	resp1.Body.Close()
	if resp1.StatusCode != http.StatusOK {
		t.Fatalf("first GET status %d; want 200", resp1.StatusCode)
	}

	// Rotate the array's session out from under us. Next call's
	// cached token is now stale -> array returns 401 -> Transport
	// re-logs-in -> retry succeeds.
	f.rotateSession("sess-2")
	loginsBefore := f.loginCalls.Load()
	dataBefore := f.dataCalls.Load()

	resp2, err := httpClient.Get(f.server.URL + "/api/2.26/volumes")
	if err != nil {
		t.Fatalf("second GET: %v", err)
	}
	resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		t.Errorf("second GET status %d; want 200 (refresh-and-retry)", resp2.StatusCode)
	}
	if got := f.loginCalls.Load() - loginsBefore; got != 1 {
		t.Errorf("login calls during retry: %d; want 1", got)
	}
	// Data call hit twice: original 401, then retry 200.
	if got := f.dataCalls.Load() - dataBefore; got != 2 {
		t.Errorf("data calls during retry: %d; want 2", got)
	}
}

func TestTransport_NonReplayableBodySurfaces401(t *testing.T) {
	f := newFakeArray(t, "api-tok-good", "sess-1")
	s := &SessionTokenSource{
		Endpoint:   f.server.URL,
		APIToken:   "api-tok-good",
		HTTPClient: f.server.Client(),
	}
	// Pre-warm with the right token, then rotate so the call gets 401.
	if _, err := s.Token(context.Background()); err != nil {
		t.Fatalf("pre-warm: %v", err)
	}
	f.rotateSession("sess-2")

	tr := &Transport{Base: f.server.Client().Transport, Source: s}
	httpClient := &http.Client{Transport: tr}

	// Hand-built request with a Body but no GetBody - the guard
	// should surface the 401 rather than retry with empty body.
	body := io.NopCloser(strings.NewReader("payload"))
	req, err := http.NewRequest(http.MethodPost, f.server.URL+"/api/2.26/volumes", body)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.GetBody = nil // simulate hand-built request that did not provide one

	loginsBefore := f.loginCalls.Load()
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("status %d; want 401 surfaced for non-replayable body", resp.StatusCode)
	}
	if got := f.loginCalls.Load() - loginsBefore; got != 0 {
		t.Errorf("login calls: %d; want 0 (no refresh attempted)", got)
	}
}

func TestTransport_NilSourceErrors(t *testing.T) {
	tr := &Transport{Base: http.DefaultTransport, Source: nil}
	req, _ := http.NewRequest(http.MethodGet, "http://example.invalid", nil)
	_, err := tr.RoundTrip(req)
	if err == nil {
		t.Fatal("expected error with nil Source; got nil")
	}
}

// TestTransport_ConcurrentRefreshSingleLogin asserts that when many
// requests share the same stale token and all observe a 401, only one
// login is performed even if the array invalidates prior sessions on
// every new login. The rest pick up the freshly-installed token via
// RefreshIfStale and retry successfully.
func TestTransport_ConcurrentRefreshSingleLogin(t *testing.T) {
	f := newFakeArray(t, "api-tok-good", "sess-initial")
	f.rotateOnLogin = true
	s := &SessionTokenSource{
		Endpoint:   f.server.URL,
		APIToken:   "api-tok-good",
		HTTPClient: f.server.Client(),
	}
	// Pre-warm. The login rotates the server to sess-rotated-1
	// (because rotateOnLogin is set) and caches it locally.
	if _, err := s.Token(context.Background()); err != nil {
		t.Fatalf("pre-warm: %v", err)
	}
	// Force the server's session away from what the cache holds, so
	// every subsequent request initially sees a 401.
	f.rotateSession("sess-evicted-out-of-band")

	tr := &Transport{Base: f.server.Client().Transport, Source: s}
	httpClient := &http.Client{Transport: tr}

	loginsBefore := f.loginCalls.Load()

	const N = 50
	var (
		wg           sync.WaitGroup
		ready, start sync.WaitGroup
		statuses     = make([]int, N)
		errs         = make([]error, N)
	)
	ready.Add(N)
	start.Add(1)

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ready.Done()
			start.Wait() // release all goroutines at once
			resp, err := httpClient.Get(f.server.URL + "/api/2.26/volumes")
			if err != nil {
				errs[i] = err
				return
			}
			statuses[i] = resp.StatusCode
			resp.Body.Close()
		}(i)
	}
	ready.Wait()
	start.Done()
	wg.Wait()

	for i := 0; i < N; i++ {
		if errs[i] != nil {
			t.Errorf("g%d: %v", i, errs[i])
			continue
		}
		if statuses[i] != http.StatusOK {
			t.Errorf("g%d status %d; want 200", i, statuses[i])
		}
	}
	if logins := f.loginCalls.Load() - loginsBefore; logins != 1 {
		t.Errorf("concurrent 401s caused %d logins; want exactly 1 (RefreshIfStale singleflight)", logins)
	}
}

// TestTransport_ReplayableBodyRetriedOn401 asserts that when a POST
// with a replayable body (GetBody set) gets a 401, the retry carries
// the same bytes as the original request - not an empty/consumed body.
func TestTransport_ReplayableBodyRetriedOn401(t *testing.T) {
	f := newFakeArray(t, "api-tok-good", "sess-1")
	s := &SessionTokenSource{
		Endpoint:   f.server.URL,
		APIToken:   "api-tok-good",
		HTTPClient: f.server.Client(),
	}
	if _, err := s.Token(context.Background()); err != nil {
		t.Fatalf("pre-warm: %v", err)
	}
	f.rotateSession("sess-2")

	tr := &Transport{Base: f.server.Client().Transport, Source: s}
	httpClient := &http.Client{Transport: tr}

	const payload = `{"name":"vol01"}`
	req, err := http.NewRequest(http.MethodPost,
		f.server.URL+"/api/2.26/volumes", strings.NewReader(payload))
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	// http.NewRequest auto-populates GetBody for strings.Reader
	// bodies; this is the contract Transport relies on for retry.
	if req.GetBody == nil {
		t.Fatal("NewRequest did not set GetBody for strings.Reader body")
	}

	f.bodiesMu.Lock()
	bodiesBefore := len(f.dataBodies)
	f.bodiesMu.Unlock()

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status %d; want 200 (retry after refresh)", resp.StatusCode)
	}

	f.bodiesMu.Lock()
	got := append([][]byte(nil), f.dataBodies[bodiesBefore:]...)
	f.bodiesMu.Unlock()
	if len(got) != 2 {
		t.Fatalf("captured %d data-request bodies; want 2 (initial 401 + retry)", len(got))
	}
	for i, b := range got {
		if string(b) != payload {
			t.Errorf("data body #%d = %q; want %q (must be replayed verbatim on retry)", i, string(b), payload)
		}
	}
}

// TestSessionTokenSource_EmptyXAuthTokenIsError asserts that a login
// response with 200 OK but a missing/empty X-Auth-Token header is
// treated as a failure - not silently cached as the empty string,
// which would poison every subsequent data call.
func TestSessionTokenSource_EmptyXAuthTokenIsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.26/login" && r.Method == http.MethodPost {
			// Deliberately do not set X-Auth-Token.
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(srv.Close)

	s := &SessionTokenSource{
		Endpoint:   srv.URL,
		APIToken:   "api-tok",
		HTTPClient: srv.Client(),
	}
	_, err := s.Token(context.Background())
	if err == nil {
		t.Fatal("expected error when login returns no X-Auth-Token; got nil")
	}
	if !strings.Contains(err.Error(), "X-Auth-Token") {
		t.Errorf("error %q does not mention X-Auth-Token", err.Error())
	}
	// The empty token must not be cached. Direct unexported-field
	// access - same package.
	s.mu.Lock()
	cached := s.session
	s.mu.Unlock()
	if cached != "" {
		t.Errorf("session cached as %q after failed login; want empty (no caching of broken state)", cached)
	}
}

// TestSessionTokenSource_ContextCancelDuringLogin asserts that a
// context cancellation while a login is in flight surfaces a
// ctx-related error and does not poison the cache.
func TestSessionTokenSource_ContextCancelDuringLogin(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Block until the request's context is cancelled, then
		// return - no headers, no body, no point pretending to
		// answer because the caller has already given up.
		<-r.Context().Done()
	}))
	t.Cleanup(srv.Close)

	s := &SessionTokenSource{
		Endpoint:   srv.URL,
		APIToken:   "api-tok",
		HTTPClient: srv.Client(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	_, err := s.Token(ctx)
	if err == nil {
		t.Fatal("expected ctx-related error from Token; got nil")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("error %v: errors.Is(err, context.DeadlineExceeded) is false", err)
	}
	s.mu.Lock()
	cached := s.session
	s.mu.Unlock()
	if cached != "" {
		t.Errorf("session cached as %q after cancelled login; want empty", cached)
	}
}

// erroringTokenSource is a TokenSource that hands out a fixed initial
// token and always fails RefreshIfStale with a configurable error.
// Used to drive Transport down its 401-then-refresh-fails path
// without needing to mutate the fake array's credentials at runtime.
type erroringTokenSource struct {
	initial      string
	refreshErr   error
	refreshCalls atomic.Int32
}

func (s *erroringTokenSource) Token(_ context.Context) (string, error) {
	return s.initial, nil
}

func (s *erroringTokenSource) RefreshIfStale(_ context.Context, _ string) (string, error) {
	s.refreshCalls.Add(1)
	return "", s.refreshErr
}

// TestTransport_RefreshFailureAfterDataChallengeSurfacesError asserts
// that when a data call gets 401 and the subsequent token refresh
// itself fails, Transport surfaces the refresh error rather than
// silently returning the original 401 response or retrying forever.
//
// Uses an injected TokenSource implementation rather than mutating
// the fake array's credentials, exercising the TokenSource interface
// in the bargain.
func TestTransport_RefreshFailureAfterDataChallengeSurfacesError(t *testing.T) {
	f := newFakeArray(t, "api-tok-real", "sess-real")
	src := &erroringTokenSource{
		initial:    "wrong-token-will-401",
		refreshErr: errors.New("simulated refresh failure"),
	}
	tr := &Transport{Base: f.server.Client().Transport, Source: src}
	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Get(f.server.URL + "/api/2.26/volumes")
	if err == nil {
		resp.Body.Close()
		t.Fatal("expected error when refresh fails after 401; got nil")
	}
	if !strings.Contains(err.Error(), "simulated refresh failure") {
		t.Errorf("error %q does not contain refresh-error text", err.Error())
	}
	if got := src.refreshCalls.Load(); got != 1 {
		t.Errorf("RefreshIfStale called %d times; want exactly 1 (no looping)", got)
	}
}

// TestTransport_RefreshFailurePreservesOriginal401Body asserts that
// when the array's data-call 401 carries a JSON error body and the
// subsequent refresh fails, the original 401 body is preserved in
// the returned error - not silently dropped. The body is the most
// useful diagnostic the array gave us; losing it leaves the caller
// with only the (often generic) refresh-failure text.
func TestTransport_RefreshFailurePreservesOriginal401Body(t *testing.T) {
	const arrayMessage = `{"errors":[{"message":"User has been disabled"}]}`
	f := newFakeArray(t, "api-tok-real", "sess-real")
	f.data401Body = []byte(arrayMessage)

	src := &erroringTokenSource{
		initial:    "wrong-token-will-401",
		refreshErr: errors.New("simulated refresh failure"),
	}
	tr := &Transport{Base: f.server.Client().Transport, Source: src}
	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Get(f.server.URL + "/api/2.26/volumes")
	if err == nil {
		resp.Body.Close()
		t.Fatal("expected error when refresh fails after 401; got nil")
	}
	if !strings.Contains(err.Error(), "simulated refresh failure") {
		t.Errorf("error %q does not contain refresh-error text", err.Error())
	}
	if !strings.Contains(err.Error(), "User has been disabled") {
		t.Errorf("error %q does not preserve the original 401 body content", err.Error())
	}
}
