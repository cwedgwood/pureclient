VERSION ?=

# Generator version. Bumping this in lockstep with regen is the
# whole point of pinning it. generate.sh has the same value as a
# fallback; if you change one, change both.
OAPI_CODEGEN_VERSION ?= v2.7.0

.PHONY: default generate install-tools clean

default:
	go vet ./...
	go build ./...
	go test ./...

generate: install-tools
	./generate.sh $(VERSION)

install-tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)

clean:
	rm -f *~ */*~
	rm -f listvol listvol-session createvol
	go clean ./...
