package pureclient_test

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

const flashArrayAPIVersion = "2.26"

type openAPISpec struct {
	Info struct {
		Version string `yaml:"version"`
	} `yaml:"info"`
	Paths map[string]any `yaml:"paths"`
}

func validateFlashArrayAPIVersion(data []byte) error {
	var spec openAPISpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return fmt.Errorf("parse OpenAPI document: %w", err)
	}
	if spec.Info.Version != flashArrayAPIVersion {
		return fmt.Errorf("info.version is %q, want %q", spec.Info.Version, flashArrayAPIVersion)
	}
	if len(spec.Paths) == 0 {
		return fmt.Errorf("OpenAPI document has no paths")
	}

	wantPrefix := "/api/" + flashArrayAPIVersion + "/"
	var wrongPaths []string
	versionedPaths := 0
	for path := range spec.Paths {
		if path == "/api/api_version" {
			continue
		}
		if !strings.HasPrefix(path, "/api/") {
			continue
		}
		versionedPaths++
		if !strings.HasPrefix(path, wantPrefix) {
			wrongPaths = append(wrongPaths, path)
		}
	}
	if versionedPaths == 0 {
		return fmt.Errorf("OpenAPI document has no versioned FlashArray API paths")
	}
	if len(wrongPaths) != 0 {
		sort.Strings(wrongPaths)
		return fmt.Errorf("paths do not target FlashArray API %s: %s",
			flashArrayAPIVersion, strings.Join(wrongPaths, ", "))
	}
	return nil
}

func TestCommittedSpecTargetsFlashArray226(t *testing.T) {
	data, err := os.ReadFile("client/api/openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if err := validateFlashArrayAPIVersion(data); err != nil {
		t.Fatal(err)
	}
}

func TestValidateFlashArrayAPIVersion(t *testing.T) {
	tests := []struct {
		name    string
		spec    string
		wantErr string
	}{
		{
			name: "valid",
			spec: `info:
  version: "2.26"
paths:
  /api/api_version: {}
  /api/2.26/volumes: {}
  /oauth2/1.0/token: {}
`,
		},
		{
			name: "wrong declared version",
			spec: `info:
  version: "2.27"
paths:
  /api/2.26/volumes: {}
`,
			wantErr: `info.version is "2.27", want "2.26"`,
		},
		{
			name: "mixed endpoint versions",
			spec: `info:
  version: "2.26"
paths:
  /api/2.26/volumes: {}
  /api/2.25/hosts: {}
`,
			wantErr: "paths do not target FlashArray API 2.26: /api/2.25/hosts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateFlashArrayAPIVersion([]byte(tt.spec))
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("validateFlashArrayAPIVersion() error = %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("validateFlashArrayAPIVersion() error = %v, want containing %q", err, tt.wantErr)
			}
		})
	}
}
