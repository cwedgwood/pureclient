# createvol example

"Create a volume" example demonstrating the write side of the API:
POST with a typed JSON request body, parameter struct for query
inputs, and status-aware response handling for the common success /
conflict / auth-error outcomes.

Uses the [`auth`](../../auth/) package for session-token caching and
transparent re-login on 401 - the same construction as
[`../listvol-session/`](../listvol-session/).

## Run

```sh
export PURE_ENDPOINT=https://array-vip
export PURE_API_TOKEN=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
export VOLUME_NAME=demo01
export VOLUME_SIZE_GIB=1
go run ./examples/createvol
```

TLS verification is hardcoded off for lab convenience (same as the
other examples). Production users should remove
`InsecureSkipVerify` and point at a real CA bundle.

## What it shows

- Constructing a `PostApi226VolumesJSONRequestBody` (= `VolumePost`)
  with required and optional fields. `Provisioned` is `*int64`, so
  the caller must take the address of the value.
- Splitting inputs between `PostApi226VolumesParams` (query / header
  parameters like the target volume name(s)) and the JSON body
  (resource attributes). oapi-codegen derives this split directly
  from the spec's parameter declarations.
- Inspecting `resp.JSON200` for the success path and
  `resp.StatusCode()` + `resp.Body` for non-2xx. The same response
  shape carries the typed success payload and the raw bytes for
  error reporting.
- Translating GiB to bytes at the boundary, because the API takes
  raw bytes for size fields.

## Expected outcomes

| Outcome | Status | Where to look |
|---|---|---|
| Success | 200 | `resp.JSON200.Items[0]` - the created Volume |
| Name in use | 409 | `resp.Body` carries `{"errors":[{"message":"Volume already exists."}]}` |
| Invalid size / params | 400 | `resp.Body` carries the array's validation error |
| Bad API token | (login error) | Initial login fails before any data call; the error message contains the array's login response. The auth helper does not retry login failures. |

## What it does NOT show

- Cleanup. The created volume persists on the array; delete it with
  `DeleteApi226VolumesWithResponse` (or `purearray volume destroy`)
  if you don't want it around.
- Volumes with QoS / priority / source (copy). Those are nested
  `allOf` fields emitted as `map[string]interface{}` - see
  [`../../RATIONALE.md`](../../RATIONALE.md) for the trade-off
  and [`../../MIGRATION.md`](../../MIGRATION.md) for assembly.
- Batch creation (POST accepts multiple names; this example only
  uses one).
