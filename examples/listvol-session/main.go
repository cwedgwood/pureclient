// Command listvol-session prints a FlashArray's volumes using the
// pureclient/auth helper for session-token caching and transparent
// re-login on 401.
//
// The auth package exposes two pieces shaped after golang.org/x/oauth2:
//
//   - auth.SessionTokenSource caches the x-auth-token and refreshes it
//     by re-running POST /api/X.Y/login with the long-lived API token.
//   - auth.Transport wraps an http.RoundTripper to inject the session
//     token on every request and retry once on a 401.
//
// Compare to examples/listvol for the minimal one-shot version that
// does not handle session expiry.
//
// Usage:
//
//	PURE_ENDPOINT=https://array-vip PURE_API_TOKEN=xxx go run ./examples/listvol-session
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/cwedgwood/pureclient/auth"
	pureclient "github.com/cwedgwood/pureclient/client"
)

func main() {
	endpoint := os.Getenv("PURE_ENDPOINT")
	apiToken := os.Getenv("PURE_API_TOKEN")
	if endpoint == "" || apiToken == "" {
		log.Fatal("set PURE_ENDPOINT (e.g. https://array-vip) and PURE_API_TOKEN")
	}

	// Lab convenience: arrays present self-signed certs by default.
	// For production use, point at a real CA bundle.
	baseTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	baseHTTPClient := &http.Client{Transport: baseTransport}

	ctx := context.Background()

	source := &auth.SessionTokenSource{
		Endpoint:   endpoint,
		APIToken:   apiToken,
		HTTPClient: baseHTTPClient,
	}
	authedHTTPClient := &http.Client{
		Transport: &auth.Transport{Base: baseTransport, Source: source},
	}

	client, err := pureclient.NewClientWithResponses(endpoint,
		pureclient.WithHTTPClient(authedHTTPClient),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.GetApi226VolumesWithResponse(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("list volumes: %s: %s", resp.Status(), string(resp.Body))
	}
	if resp.JSON200 == nil || resp.JSON200.Items == nil {
		return
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(*resp.JSON200.Items); err != nil {
		log.Fatal(err)
	}
}
