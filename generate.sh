#!/bin/bash
set -euo pipefail

# Regenerate the Go client from an Everpure (formerly Pure Storage) FlashArray spec.
#
# The committed spec at client/api/openapi.yaml is OpenAPI 3.0.1, originally
# converted from the upstream Swagger 2.0 spec by openapi-generator. The
# version-extraction mode below only supports OAS3 specs (FA2.42 and newer);
# older versions are Swagger 2.0 and oapi-codegen will not parse them.
#
# Prerequisites:
#   - oapi-codegen v2.7.0 (run `make install-tools`, or
#     `go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.0`)
#   - Docker (only needed when extracting a spec by version)
#   - Go toolchain
#
# Usage:
#   ./generate.sh                  # use existing spec in client/api/openapi.yaml
#   ./generate.sh /path/to.yaml    # use a local spec file (does NOT update openapi.yaml)
#   ./generate.sh FA2.42           # extract from quay.io/purestorage/swagger image (OAS3 only)
#                                  # AND replace client/api/openapi.yaml with the extracted spec
#                                  # so the committed spec matches the committed client.
#
# Note: the version-extraction form (./generate.sh FA2.NN) only works for
# OAS3 versions (FA2.42 and later). Earlier versions including the
# currently-committed FA2.26 are Swagger 2.0 and require a one-shot
# Swagger 2.0 -> OpenAPI 3.0.1 conversion that this script does not
# automate. To regenerate against FA2.26 specifically, re-use the
# already-converted client/api/openapi.yaml (the default no-arg form).

SPEC_ARG="${1:-}"
REPO_ROOT="$(cd "$(dirname "$0")" && pwd)"
CLIENT_DIR="${REPO_ROOT}/client"
# Pinned to a specific version tag rather than :latest so that re-runs
# pull a known image. v1.88.0 is the same digest (sha256:8d1df156...)
# that produced the committed client/api/openapi.yaml.
SWAGGER_IMAGE="quay.io/purestorage/swagger:v1.88.0"
# Generator version. Kept in sync with OAPI_CODEGEN_VERSION in the
# Makefile (run `make install-tools` to install the pinned version).
OAPI_CODEGEN_VERSION="v2.7.0"
OAPI_CODEGEN="${OAPI_CODEGEN:-oapi-codegen}"

# Find oapi-codegen
if ! command -v "${OAPI_CODEGEN}" &>/dev/null; then
    GOBIN="$(go env GOPATH)/bin"
    if [[ -x "${GOBIN}/oapi-codegen" ]]; then
        OAPI_CODEGEN="${GOBIN}/oapi-codegen"
    else
        echo "Error: oapi-codegen not found. Install with:" >&2
        echo "  make install-tools" >&2
        echo "or:" >&2
        echo "  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION}" >&2
        exit 1
    fi
fi

spec_file=""
cleanup_spec=""
extracted_source_path=""
trap 'if [[ -n "${cleanup_spec}" && -f "${cleanup_spec}" ]]; then rm -f "${cleanup_spec}"; fi' EXIT

if [[ -z "${SPEC_ARG}" ]]; then
    spec_file="${CLIENT_DIR}/api/openapi.yaml"
    if [[ ! -f "${spec_file}" ]]; then
        echo "Error: no spec found at ${spec_file}" >&2
        echo "Provide a spec file or version as argument." >&2
        exit 1
    fi
    echo "Using existing spec: ${spec_file}"
elif [[ -f "${SPEC_ARG}" ]]; then
    spec_file="$(realpath "${SPEC_ARG}")"
    echo "Using local spec: ${spec_file}"
    echo "Note: ${CLIENT_DIR}/api/openapi.yaml will NOT be updated; the regenerated" >&2
    echo "      client may not match the committed spec. Either update openapi.yaml" >&2
    echo "      by hand or use the './generate.sh FA<version>' form for known-good provenance." >&2
else
    spec_file="${REPO_ROOT}/${SPEC_ARG}.spec.yaml"
    cleanup_spec="${spec_file}"
    extracted_source_path="/usr/share/pureswagger/html/specs/${SPEC_ARG}.spec.yaml"
    echo "Extracting ${SPEC_ARG}.spec.yaml from ${SWAGGER_IMAGE}..."
    docker run --rm "${SWAGGER_IMAGE}" \
        cat "${extracted_source_path}" > "${spec_file}"
    if [[ ! -s "${spec_file}" ]]; then
        echo "Error: spec ${SPEC_ARG}.spec.yaml not found in image." >&2
        echo "Run: docker run --rm ${SWAGGER_IMAGE} ls /usr/share/pureswagger/html/specs/" >&2
        exit 1
    fi
    if head -1 "${spec_file}" | grep -q '^swagger:'; then
        echo "Error: ${SPEC_ARG}.spec.yaml is Swagger 2.0; oapi-codegen requires OpenAPI 3." >&2
        echo "Older versions (FA2.41 and earlier) are Swagger 2.0 and must be converted first," >&2
        echo "e.g. with 'openapi-generator generate -i <spec> -g openapi-yaml'. Try a newer version" >&2
        echo "(FA2.42+) for native OAS3, or use the committed client/api/openapi.yaml." >&2
        exit 1
    fi
    echo "Extracted: ${spec_file}"
fi

echo ""
echo "Generating types..."
"${OAPI_CODEGEN}" --package client --generate types \
    -o "${CLIENT_DIR}/types.gen.go" "${spec_file}"

echo "Generating client..."
"${OAPI_CODEGEN}" --package client --generate client \
    -o "${CLIENT_DIR}/client.gen.go" "${spec_file}"

# Extraction mode: publish the extracted spec to client/api/openapi.yaml
# with a freshly-generated provenance header, so the committed spec
# matches the committed client. Disarm the cleanup trap since the
# tempfile is now the canonical spec.
if [[ -n "${extracted_source_path}" ]]; then
    echo ""
    echo "Publishing extracted spec to ${CLIENT_DIR}/api/openapi.yaml..."

    image_digest="$(docker inspect --format='{{index .RepoDigests 0}}' "${SWAGGER_IMAGE}" 2>/dev/null | sed 's/.*@//' || true)"
    if [[ -z "${image_digest}" ]]; then
        image_digest="(digest unavailable; rerun 'docker pull ${SWAGGER_IMAGE}' to populate)"
    fi
    generated_at="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
    spec_with_header="${spec_file}.with-header"
    {
        printf '%s\n' "# Provenance:"
        printf '%s\n' "#   Upstream:    ${SWAGGER_IMAGE}"
        printf '%s\n' "#                (${image_digest})"
        printf '%s\n' "#   Source file: ${extracted_source_path} (OpenAPI 3.0.x; no conversion needed)"
        printf '%s\n' "#   Consumer:    github.com/oapi-codegen/oapi-codegen ${OAPI_CODEGEN_VERSION}"
        printf '%s\n' "#   Generated:   ${generated_at} by 'generate.sh ${SPEC_ARG}'"
        printf '%s\n' "# See generate.sh, README.md, and MIGRATION.md."
        cat "${spec_file}"
    } > "${spec_with_header}"
    mv "${spec_with_header}" "${CLIENT_DIR}/api/openapi.yaml"
    rm -f "${spec_file}"
    cleanup_spec=""  # disarm trap; spec has been published, not deleted
    echo "Updated: ${CLIENT_DIR}/api/openapi.yaml (commit it alongside the regenerated client)"
fi

echo ""
echo "Running go mod tidy..."
cd "${REPO_ROOT}"
go mod tidy

echo ""
echo "Verifying build..."
go vet ./...
go build ./...
go test ./auth/

echo ""
echo "Done. Client regenerated from $(basename "${spec_file}")."
echo "Review changes with: git diff --stat"
