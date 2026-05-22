package auth_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cwedgwood/pureclient/auth"
	pureclient "github.com/cwedgwood/pureclient/client"
)

// fakeArrayServer stands in for a real FlashArray in these examples.
// A real caller would point Endpoint at the array VIP and use the
// array's TLS-aware http.Client.
func fakeArrayServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.26/login" && r.Method == http.MethodPost {
			w.Header().Set("X-Auth-Token", "session-token-xyz")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"items":[]}`))
	}))
}

// Wiring the auth helpers into a pureclient.ClientWithResponses.
// Transport handles session-token injection on every request and
// transparently re-logs-in + retries once on a 401.
func Example() {
	srv := fakeArrayServer()
	defer srv.Close()

	source := &auth.SessionTokenSource{
		Endpoint:   srv.URL,
		APIToken:   "long-lived-api-token",
		HTTPClient: srv.Client(), // base client for the login call itself
	}
	httpClient := &http.Client{
		Transport: &auth.Transport{
			Base:   srv.Client().Transport,
			Source: source,
		},
	}
	client, err := pureclient.NewClientWithResponses(srv.URL,
		pureclient.WithHTTPClient(httpClient))
	if err != nil {
		fmt.Println("construct:", err)
		return
	}
	resp, err := client.GetApi226VolumesWithResponse(context.Background(), nil)
	if err != nil {
		fmt.Println("call:", err)
		return
	}
	fmt.Println("status:", resp.StatusCode())
	// Output: status: 200
}

// SessionTokenSource can be used standalone (without Transport) when
// the caller wants explicit control over when logins happen, for
// example to pre-warm the cache before serving the first request.
func ExampleSessionTokenSource_Token() {
	srv := fakeArrayServer()
	defer srv.Close()

	source := &auth.SessionTokenSource{
		Endpoint:   srv.URL,
		APIToken:   "long-lived-api-token",
		HTTPClient: srv.Client(),
	}
	token, err := source.Token(context.Background())
	if err != nil {
		fmt.Println("token:", err)
		return
	}
	fmt.Println("got token:", token)
	// Output: got token: session-token-xyz
}

// Refresh forces a fresh login regardless of whether the cached
// token is still valid. Use it to recover from out-of-band session
// invalidation (array failover, session-store restart, ...) rather
// than waiting for the next request to hit a 401.
func ExampleSessionTokenSource_Refresh() {
	srv := fakeArrayServer()
	defer srv.Close()

	source := &auth.SessionTokenSource{
		Endpoint:   srv.URL,
		APIToken:   "long-lived-api-token",
		HTTPClient: srv.Client(),
	}
	if err := source.Refresh(context.Background()); err != nil {
		fmt.Println("refresh:", err)
		return
	}
	fmt.Println("session refreshed")
	// Output: session refreshed
}
