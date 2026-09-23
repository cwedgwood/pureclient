package tlsconfig

import (
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestDefaultRejectsUntrustedCertificate(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()

	t.Setenv(CAFileEnv, "")
	t.Setenv(InsecureSkipVerifyEnv, "")
	client := newClient(t)

	if _, err := client.Get(server.URL); err == nil {
		t.Fatal("default TLS configuration accepted an untrusted certificate")
	}
}

func TestConfiguredCASucceeds(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	caFile := writeCertificate(t, server)
	t.Setenv(CAFileEnv, caFile)
	t.Setenv(InsecureSkipVerifyEnv, "")
	client := newClient(t)

	response, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("GET with configured CA: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %s, want 204 No Content", response.Status)
	}
}

func TestInsecureRequiresExplicitOptIn(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()

	t.Setenv(CAFileEnv, "")
	t.Setenv(InsecureSkipVerifyEnv, "true")
	transport, insecure, err := NewTransport()
	if err != nil {
		t.Fatalf("NewTransport: %v", err)
	}
	if !insecure {
		t.Fatal("insecure = false, want true")
	}
	response, err := (&http.Client{Transport: transport}).Get(server.URL)
	if err != nil {
		t.Fatalf("GET with explicit insecure opt-in: %v", err)
	}
	response.Body.Close()
}

func TestInvalidConfiguration(t *testing.T) {
	tests := []struct {
		name        string
		insecure    string
		caFile      string
		wantInError string
	}{
		{
			name:        "malformed insecure value",
			insecure:    "yes",
			wantInError: `PURE_INSECURE_SKIP_VERIFY must be "true" or "false", got "yes"`,
		},
		{
			name:        "unreadable CA file",
			caFile:      "does-not-exist.pem",
			wantInError: `read PURE_CA_FILE "does-not-exist.pem"`,
		},
		{
			name:        "invalid CA file",
			caFile:      writeInvalidCA(t),
			wantInError: "no PEM-encoded certificates found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(InsecureSkipVerifyEnv, tt.insecure)
			t.Setenv(CAFileEnv, tt.caFile)
			if _, _, err := NewTransport(); err == nil {
				t.Fatal("NewTransport succeeded, want error")
			} else if !strings.Contains(err.Error(), tt.wantInError) {
				t.Fatalf("error = %q, want it to contain %q", err, tt.wantInError)
			}
		})
	}
}

func newClient(t *testing.T) *http.Client {
	t.Helper()
	transport, insecure, err := NewTransport()
	if err != nil {
		t.Fatalf("NewTransport: %v", err)
	}
	if insecure {
		t.Fatal("insecure = true, want false")
	}
	return &http.Client{Transport: transport}
}

func writeCertificate(t *testing.T, server *httptest.Server) string {
	t.Helper()
	certificate := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: server.Certificate().Raw,
	})
	return writeTestFile(t, "test-ca.pem", certificate)
}

func writeInvalidCA(t *testing.T) string {
	t.Helper()
	return writeTestFile(t, "test-invalid-ca.pem", []byte("not a certificate\n"))
}

func writeTestFile(t *testing.T, name string, content []byte) string {
	t.Helper()
	path := fmt.Sprintf(".%s-%d", name, os.Getpid())
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			t.Errorf("remove %s: %v", path, err)
		}
	})
	return path
}
