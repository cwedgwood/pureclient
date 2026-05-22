// Command createvol creates a FlashArray volume.
//
// Demonstrates the write side of the API: POST with a JSON request
// body, parameters via a typed Params struct, and status-aware
// response handling for the common success / conflict / auth-error
// outcomes.
//
// Uses the pureclient/auth helper for session-token caching and
// transparent re-login on 401; the same construction shape works
// for any write call.
//
// Usage:
//
//	PURE_ENDPOINT=https://array-vip \
//	PURE_API_TOKEN=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
//	VOLUME_NAME=demo01 \
//	VOLUME_SIZE_GIB=1 \
//	go run ./examples/createvol
//
// Reports the created Volume on success, or the array's error
// message on a non-2xx (typically 409 if the volume name already
// exists, 400 if size is invalid).
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cwedgwood/pureclient/auth"
	pureclient "github.com/cwedgwood/pureclient/client"
)

func main() {
	endpoint := os.Getenv("PURE_ENDPOINT")
	apiToken := os.Getenv("PURE_API_TOKEN")
	name := os.Getenv("VOLUME_NAME")
	sizeGiBStr := os.Getenv("VOLUME_SIZE_GIB")
	if endpoint == "" || apiToken == "" || name == "" || sizeGiBStr == "" {
		log.Fatal("set PURE_ENDPOINT, PURE_API_TOKEN, VOLUME_NAME, VOLUME_SIZE_GIB")
	}
	sizeGiB, err := strconv.ParseInt(sizeGiBStr, 10, 64)
	if err != nil || sizeGiB <= 0 {
		log.Fatalf("VOLUME_SIZE_GIB must be a positive integer, got %q", sizeGiBStr)
	}
	sizeBytes := sizeGiB * 1024 * 1024 * 1024

	// Lab convenience: arrays present self-signed certs by default.
	// For production use, point at a real CA bundle.
	baseTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	baseHTTPClient := &http.Client{Transport: baseTransport}

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

	// Params carry query/header inputs (here: the volume name(s) on
	// which to operate). Body carries the request payload (here:
	// the volume's desired attributes). Both are split out by
	// oapi-codegen from the spec's parameter declaration.
	names := []string{name}
	params := &pureclient.PostApi226VolumesParams{Names: &names}
	body := pureclient.PostApi226VolumesJSONRequestBody{
		Provisioned: &sizeBytes,
	}

	resp, err := client.PostApi226VolumesWithResponse(context.Background(), params, body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("create volume %q: %s: %s", name, resp.Status(), string(resp.Body))
	}
	if resp.JSON200 == nil || resp.JSON200.Items == nil || len(*resp.JSON200.Items) == 0 {
		log.Fatalf("create volume %q: 200 OK but empty response body", name)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode((*resp.JSON200.Items)[0]); err != nil {
		log.Fatal(err)
	}
}
