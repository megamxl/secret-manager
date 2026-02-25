package authentication

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"secret-manager/internal/config"
)

// MTLSMiddleware rejects requests that did not present a verified client cert.
// The actual TLS handshake must be configured with tls.RequireAndVerifyClientCert.
func MTLSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil || len(r.TLS.VerifiedChains) == 0 {
			http.Error(w, "mTLS: client certificate required", http.StatusUnauthorized)
			return
		}

		leaf := r.TLS.VerifiedChains[0][0]
		id := Identity{
			Method:  AuthMethodMTLS,
			Subject: leaf.Subject.String(),
		}

		next.ServeHTTP(w, WithIdentity(r, id))
	})
}

// BuildTLSConfig builds a *tls.Config for the server from the config.
func BuildTLSConfig(cfg config.TLSConfig) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("tls: load keypair: %w", err)
	}

	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	if cfg.ClientCAFile != "" {
		cas, err := os.ReadFile(cfg.ClientCAFile)
		if err != nil {
			return nil, fmt.Errorf("tls: read client CA: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(cas) {
			return nil, fmt.Errorf("tls: no valid certs in client CA file")
		}
		tlsCfg.ClientCAs = pool
		tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
	}

	return tlsCfg, nil
}
