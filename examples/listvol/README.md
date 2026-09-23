# listvol example

Minimal "list volumes" example demonstrating the one-shot login
pattern: POST `/api/2.26/login` once at startup, capture the
`X-Auth-Token` from the response header, inject it on every
subsequent request via `pureclient.WithRequestEditorFn`.

## Run

```sh
export PURE_ENDPOINT=https://array-vip
export PURE_API_TOKEN=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
go run ./examples/listvol
```

TLS certificates are verified using the system trust roots by default.
For an array using a private CA, set `PURE_CA_FILE` to a PEM certificate
bundle; its certificates are added to, rather than replacing, the system
roots. As a last resort for a lab, set
`PURE_INSECURE_SKIP_VERIFY=true`; the command prints a warning when
verification is disabled. Only the exact values `true` and `false` are
accepted.

## What it shows

- Constructing `*ClientWithResponses` with `WithHTTPClient` and
  `WithRequestEditorFn`.
- Logging in via `PostApi226LoginWithResponse` and reading the
  session token from `resp.HTTPResponse.Header.Get("X-Auth-Token")`.
- Listing volumes via `GetApi226VolumesWithResponse`, walking the
  pointer-to-slice `Items` field, and dereferencing pointer string
  fields safely.
- Surfacing non-2xx responses by inspecting `resp.StatusCode()` and
  the raw `resp.Body`.

## What it does NOT show

- Session expiry recovery. If the session token is invalidated
  between login and the volume list call, this example will fail
  with a 401. See [`../listvol-session/`](../listvol-session/)
  for the session-aware variant that uses the
  [`auth`](../../auth/) package to re-login transparently on 401.
- Pagination (the array returns up to a default limit per call).
- Filtering or sorting (`GetApi226VolumesParams.Filter`,
  `Sort`, `Limit`, `Offset`).
