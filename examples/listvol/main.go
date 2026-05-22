// Command listvol prints a FlashArray's volumes using a one-shot login.
//
// Logs in once at startup, captures the x-auth-token from the
// response, and injects it on every subsequent request via
// WithRequestEditorFn. Does not handle 401s (e.g. session expiry).
// See examples/listvol-session for a session-aware version that
// re-logs-in on demand.
//
// Usage:
//
//	PURE_ENDPOINT=https://array-vip PURE_API_TOKEN=xxx go run ./examples/listvol
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	pureclient "github.com/cwedgwood/pureclient/client"
)

// login does a one-shot POST /api/2.26/login using the API token and
// returns the x-auth-token session token from the response header.
func login(ctx context.Context, endpoint, apiToken string, httpClient *http.Client) (string, error) {
	c, err := pureclient.NewClientWithResponses(endpoint,
		pureclient.WithHTTPClient(httpClient),
	)
	if err != nil {
		return "", err
	}
	resp, err := c.PostApi226LoginWithResponse(ctx, &pureclient.PostApi226LoginParams{
		ApiToken: &apiToken,
	})
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("login failed: %s: %s", resp.Status(), string(resp.Body))
	}
	token := resp.HTTPResponse.Header.Get("X-Auth-Token")
	if token == "" {
		return "", errors.New("login succeeded but X-Auth-Token header was empty")
	}
	return token, nil
}

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

	sessionToken, err := login(ctx, endpoint, apiToken, baseHTTPClient)
	if err != nil {
		log.Fatal(err)
	}

	client, err := pureclient.NewClientWithResponses(endpoint,
		pureclient.WithHTTPClient(baseHTTPClient),
		pureclient.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("x-auth-token", sessionToken)
			return nil
		}),
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
