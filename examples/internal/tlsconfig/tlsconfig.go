// Package tlsconfig provides the TLS configuration shared by the examples.
package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

const (
	CAFileEnv             = "PURE_CA_FILE"
	InsecureSkipVerifyEnv = "PURE_INSECURE_SKIP_VERIFY"
)

// NewTransport returns a clone of the default HTTP transport configured from
// PURE_CA_FILE and PURE_INSECURE_SKIP_VERIFY.
func NewTransport() (*http.Transport, bool, error) {
	insecure, err := insecureSkipVerify()
	if err != nil {
		return nil, false, err
	}

	roots, err := rootCAs()
	if err != nil {
		return nil, false, err
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{
		MinVersion:         tls.VersionTLS12,
		RootCAs:            roots,
		InsecureSkipVerify: insecure,
	}
	return transport, insecure, nil
}

func insecureSkipVerify() (bool, error) {
	value, ok := os.LookupEnv(InsecureSkipVerifyEnv)
	if !ok || value == "" || value == "false" {
		return false, nil
	}
	if value == "true" {
		return true, nil
	}
	return false, fmt.Errorf("%s must be %q or %q, got %q", InsecureSkipVerifyEnv, "true", "false", value)
}

func rootCAs() (*x509.CertPool, error) {
	path := os.Getenv(CAFileEnv)
	if path == "" {
		return nil, nil
	}

	roots, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("load system certificate pool: %w", err)
	}
	pem, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s %q: %w", CAFileEnv, path, err)
	}
	if !roots.AppendCertsFromPEM(pem) {
		return nil, fmt.Errorf("parse %s %q: no PEM-encoded certificates found", CAFileEnv, path)
	}
	return roots, nil
}
