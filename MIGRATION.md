# Migration Guide: pureclient API Changes

This document covers the breaking changes when migrating from the old
`swagger-codegen` client (API v2.8) to the new `oapi-codegen` client
(API v2.26).

## Contents

- [Summary](#summary)
- [Client Construction](#client-construction)
- [API Calls](#api-calls)
- [Responses](#responses)
- [Model Types](#model-types)
- [Method Name Mapping](#method-name-mapping)
- [Parameters](#parameters)
- [Request Bodies](#request-bodies)
- [Dependencies](#dependencies)
- [Login Endpoint](#login-endpoint)
- [Method Discovery](#method-discovery)
- [Non-2xx Responses](#non-2xx-responses)
- [ClientWithResponses vs Client](#clientwithresponses-vs-client)
- [Custom Transports and Layered Auth](#custom-transports-and-layered-auth)

## Summary

The client was regenerated using a different code generator
([oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)) and a
newer API version. Every consumer will need code changes - this is
not a drop-in upgrade.

## Client Construction

### Before

```go
import pureclient "github.com/cwedgwood/pureclient/client"

cfg := pureclient.NewConfiguration()
cfg.BasePath = "https://array-vip:443"
cfg.DefaultHeader["x-auth-token"] = token
cfg.HTTPClient = &httpClient
client := pureclient.NewAPIClient(cfg)
```

### After

```go
import pureclient "github.com/cwedgwood/pureclient/client"

client, err := pureclient.NewClientWithResponses("https://array-vip:443",
    pureclient.WithHTTPClient(&httpClient),
    pureclient.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
        req.Header.Set("x-auth-token", token)
        return nil
    }),
)
```

**Key differences:**

- `NewConfiguration()` + `NewAPIClient()` replaced by `NewClientWithResponses()`.
- `cfg.BasePath` replaced by the first argument to `NewClientWithResponses()`.
- `cfg.DefaultHeader` replaced by `WithRequestEditorFn()` - a callback
  that can modify every outgoing request.
- `cfg.HTTPClient` replaced by `WithHTTPClient()`.
- Construction now returns `(*ClientWithResponses, error)`.

## API Calls

### Before

```go
vols, httpRes, err := client.VolumesApi.Api28VolumesGet(ctx, nil)
```

### After

```go
resp, err := client.GetApi226VolumesWithResponse(ctx, nil)
```

**Key differences:**

- Methods are on the client directly, not on sub-services
  (`client.VolumesApi.X` -> `client.X`).
- Method naming changed from `Api28<Resource><Verb>` to
  `<Verb>Api226<Resource>WithResponse` (HTTP verb moves to the front).
- API version in method names changed from `28` to `226`.
- Returns `(response, error)` instead of `(body, httpResponse, error)`.

## Responses

### Before

```go
vols, httpRes, err := client.VolumesApi.Api28VolumesGet(ctx, nil)
if err != nil { ... }
for _, v := range vols.Items {
    fmt.Println(v.Name)
}
```

### After

```go
resp, err := client.GetApi226VolumesWithResponse(ctx, nil)
if err != nil { ... }
// Check HTTP status via typed response fields
if resp.JSON200 != nil && resp.JSON200.Items != nil {
    for _, v := range *resp.JSON200.Items {
        fmt.Println(*v.Name)
    }
}
```

**Key differences:**

- The response is a struct with typed fields per HTTP status code
  (`resp.JSON200` for 200, etc.).
- `resp.HTTPResponse` gives the raw `*http.Response` if needed.
- `resp.Status()` and `resp.StatusCode()` for status inspection.
- Collection fields like `Items` are `*[]Type` (pointer to slice),
  not `[]Type`. Dereference before ranging.
- Most optional schema properties are pointer fields (`*string`,
  `*int64`, etc.) so absent can be distinguished from zero-value.
  Required fields and type aliases may be non-pointer - check the
  generated type.
- Many composed (`allOf`) nested objects are emitted as
  `map[string]interface{}` rather than typed structs (e.g.
  `Volume.Qos`, `Volume.PriorityAdjustment`, `Volume.Space`,
  `Audit.Origin`). This is a type-safety regression vs the old
  client; assemble these as map literals with the documented keys.

## Model Types

### Before

Many models had opaque names generated from the spec structure:

```go
type InlineResponse200_42 struct { ... }
type Model2_8VolumesBody struct { ... }
```

### After

All models have descriptive names derived from their schema:

```go
type VolumeGetResponse struct { ... }
type Volume struct { ... }
type VolumePatch struct { ... }
```

## Method Name Mapping

Common pattern for translating old method names to new:

| Old (v2.8, swagger-codegen) | New (v2.26, oapi-codegen) |
|---|---|
| `VolumesApi.Api28VolumesGet(ctx, nil)` | `GetApi226VolumesWithResponse(ctx, nil)` |
| `VolumesApi.Api28VolumesPost(ctx, body)` | `PostApi226VolumesWithResponse(ctx, params, body)` |
| `VolumesApi.Api28VolumesDelete(ctx, nil)` | `DeleteApi226VolumesWithResponse(ctx, nil)` |
| `HostsApi.Api28HostsGet(ctx, nil)` | `GetApi226HostsWithResponse(ctx, nil)` |
| `ArraysApi.Api28ArraysGet(ctx, nil)` | `GetApi226ArraysWithResponse(ctx, nil)` |
| `SessionsApi.Api28LoginPost(ctx, body)` | `PostApi226LoginWithResponse(ctx, params)` (see Login section) |

**Pattern:** `<Resource>Api.Api28<Resource><Verb>` ->
`<Verb>Api226<Resource>WithResponse`

## Parameters

### Before

Optional parameters were passed via a struct or `nil`:

```go
opts := &pureclient.Api28VolumesGetOpts{
    Filter: optional.NewString("name='vol1'"),
}
vols, _, err := client.VolumesApi.Api28VolumesGet(ctx, opts)
```

### After

Parameters use a typed struct (pass `nil` for defaults):

```go
filter := "name='vol1'"
params := &pureclient.GetApi226VolumesParams{
    Filter: &filter,
}
resp, err := client.GetApi226VolumesWithResponse(ctx, params)
```

**Key differences:**

- `github.com/antihax/optional` is no longer used. All optional
  parameters are pointer types - pass `nil` to omit, `&value` to set.
- Parameter structs are named `<Verb>Api226<Resource>Params`.

## Request Bodies

### Before

```go
body := pureclient.Model2_8VolumesBody{
    Provisioned: 1073741824,
}
vol, _, err := client.VolumesApi.Api28VolumesPost(ctx, body, nil)
```

### After

```go
body := pureclient.PostApi226VolumesJSONRequestBody{
    Provisioned: ptrInt64(1073741824),
}
resp, err := client.PostApi226VolumesWithResponse(ctx, params, body)
```

Body types are named `<Verb>Api226<Resource>JSONRequestBody`.

## Dependencies

| | Before | After |
|---|---|---|
| `github.com/antihax/optional` | required | removed |
| `golang.org/x/oauth2` | required | removed |
| `gopkg.in/yaml.v2` | required (example) | removed |
| `github.com/oapi-codegen/runtime` | - | required |

## Login Endpoint

The login path changed from `/api/2.8/login` to `/api/2.26/login`.
Update any hardcoded login URLs.

The generated `PostApi226LoginWithResponse` takes parameters, not a
body. The API token goes in the `api-token` header (serialized from
`PostApi226LoginParams.ApiToken`); the array returns the session
token in the `X-Auth-Token` response header.

```go
params := &pureclient.PostApi226LoginParams{ApiToken: &apiToken}
resp, err := client.PostApi226LoginWithResponse(ctx, params)
if err != nil { ... }
if resp.StatusCode() != http.StatusOK {
    log.Fatalf("login failed: %s: %s", resp.Status(), string(resp.Body))
}
sessionToken := resp.HTTPResponse.Header.Get("X-Auth-Token")
```

## Method Discovery

The pattern above is a starting point but does not always hold -
nested paths produce names like
`GetApi226ProtectionGroupsSnapshotsTransferWithResponse`, and the
auth endpoints are flat (`PostApi226LoginWithResponse`, not
`PostApi226SessionsLoginWithResponse`). To find what you need:

```sh
grep -n "WithResponse(ctx" client/client.gen.go | less
```

or `go doc github.com/cwedgwood/pureclient/client ClientWithResponsesInterface`.

## Non-2xx Responses

Generated response parsers only populate `JSON200` (or other
`JSON<code>` fields) when both the status code and content type
match. For unexpected statuses, inspect `resp.StatusCode()` and the
raw `resp.Body`:

```go
if resp.StatusCode() != http.StatusOK {
    return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
```

## ClientWithResponses vs Client

`NewClientWithResponses` returns a wrapper whose methods parse the
HTTP body into typed `JSON<code>` fields. `NewClient` returns the
lower-level client that returns raw `*http.Response`; use it when
you need to handle the response body yourself.

## Custom Transports and Layered Auth

The old generated client carried significant auth machinery that was
never used in practice - context-key helpers (`ContextOAuth2`,
`ContextBasicAuth`, `ContextAccessToken`, `ContextAPIKey`,
`ContextHttpSignatureAuth`) and a dependency on
`golang.org/x/oauth2`. The new client has **none of this**. Auth,
retries, rate-limit backoff, TLS rotation, request logging - all of
it lives below the client in a normal `http.RoundTripper`.

Wire any custom transport via `WithHTTPClient`:

```go
httpClient := &http.Client{
    Transport: &myAuthTransport{ /* ... */ },
}
client, _ := pureclient.NewClientWithResponses(endpoint,
    pureclient.WithHTTPClient(httpClient))
```

This pattern is friendly to layered behaviour. Each concern is its
own RoundTripper or its own wrapping function - the generated
client neither knows nor cares.

### Per-request header injection

For simple cases (a single static header, or a header that depends
only on per-request context), use `WithRequestEditorFn` - a
`func(context.Context, *http.Request) error` callback:

```go
pureclient.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
    req.Header.Set("x-auth-token", token)
    return nil
})
```

### Session-token refresh on 401

For the canonical FlashArray "long-lived API token -> short-lived
x-auth-token session, refresh on 401" pattern, use the
[`auth`](auth/auth.go) package shipped with this repo. It exposes
a `SessionTokenSource` (caches the session token, safe for
concurrent use) and a `Transport` (injects the cached token,
transparently re-logs-in and retries on 401, surfaces 401 on
requests with non-replayable bodies). See
[examples/listvol-session/](examples/listvol-session/) for a
runnable usage example and `go doc ./auth` (or
<https://pkg.go.dev/github.com/cwedgwood/pureclient/auth>) for the
API surface.

### Pure's OAuth2 client-credentials flow (`/oauth2/1.0/token`)

This is a real, supported FlashArray auth path - separate from the
`/api/X.Y/login` + `x-auth-token` pattern - using RFC 7523-style
JWT bearer assertion:

1. An admin registers a per-user **API Client** on the array with a
   public RSA key, a `client_id` (used as the JWT `aud`), and a
   `key_id` (used as the JWT `kid`).
2. The consumer signs a short-lived JWT with the corresponding
   private key.
3. The JWT is exchanged at `POST /oauth2/1.0/token` for a bearer
   token (typically hour-long).
4. Bearer is sent as `Authorization: Bearer <token>` on every
   subsequent request; on 401, re-exchange.

None of this changes when migrating from the old client to the new
one - the OAuth machinery is implemented as a custom
`http.RoundTripper`, plugged in via `WithHTTPClient` exactly like
the simpler session-token pattern. The new client's only relevant
change is that it no longer pulls in `golang.org/x/oauth2`
unnecessarily or invents its own context-key auth API; existing
hand-written bearer-token transports work unchanged.

The generated client is *transport-agnostic*: 429 backoff, retry,
TLS root-CA rotation, and bearer/session refresh can all be layered
into a single `http.RoundTripper` (as seen in real consumers) and
the per-call API surface stays untouched.
