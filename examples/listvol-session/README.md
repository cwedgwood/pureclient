# listvol-session example

"List volumes" example using the [`auth`](../../auth/) package for
transparent session-token caching and re-login on 401. Same task as
[`../listvol/`](../listvol/), but resilient to session expiry.

## Run

```sh
export PURE_ENDPOINT=https://array-vip
export PURE_API_TOKEN=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
go run ./examples/listvol-session
```

TLS verification is hardcoded off for lab convenience (same as the
other examples). Production users should remove `InsecureSkipVerify`
and point at a real CA bundle.

## What it shows

- Constructing an `&auth.SessionTokenSource{...}` (the cache) and
  an `&auth.Transport{...}` (the HTTP layer that injects the
  cached token and re-logs-in on 401).
- Composing the auth transport into a normal `*http.Client` and
  passing it to `pureclient.NewClientWithResponses` via
  `WithHTTPClient`.
- Letting the auth layer disappear from the call site -
  `client.GetApi226VolumesWithResponse(...)` looks identical to
  the one-shot variant; recovery from session expiry happens
  silently underneath.

## When this matters

The FlashArray session token has a short timeout (typically
30 minutes of inactivity). Long-running clients - daemons,
controllers, scheduled jobs - will outlive any single session
token and need to recover. This example is the minimum viable
shape of that recovery; the same pattern scales to clients that
make millions of calls.

## See also

- [`../listvol/`](../listvol/) - the one-shot equivalent.
- [`../../auth/`](../../auth/) - the helper package, with godoc
  examples for `SessionTokenSource.Token`,
  `SessionTokenSource.Refresh`, and the full Transport wiring.
