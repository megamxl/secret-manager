package authentication

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"secret-manager/internal/config"
	"secret-manager/internal/logging"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"go.uber.org/zap"
)

type spireMiddleware struct {
	source  *workloadapi.X509Source
	allowed map[string]struct{}
}

func NewSPIREMiddleware(cfg config.SPIREConfig) (func(http.Handler) http.Handler, error) {
	ctx := context.Background()
	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(workloadapi.WithAddr(cfg.SocketPath)),
	)
	if err != nil {
		return nil, fmt.Errorf("spire: workload API: %w", err)
	}

	allowed := make(map[string]struct{}, len(cfg.AllowedSPIFFEIDs))
	for _, id := range cfg.AllowedSPIFFEIDs {
		allowed[id] = struct{}{}
	}

	m := &spireMiddleware{source: source, allowed: allowed}
	return m.middleware, nil
}

func SPIRETLSConfig(socketPath, trustDomain string) (*tls.Config, error) {
	ctx := context.Background()
	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	)
	if err != nil {
		return nil, fmt.Errorf("spire: x509 source: %w", err)
	}

	td, err := spiffeid.TrustDomainFromString(trustDomain)
	if err != nil {
		return nil, fmt.Errorf("spire: trust domain: %w", err)
	}

	return tlsconfig.MTLSServerConfig(source, source, tlsconfig.AuthorizeMemberOf(td)), nil
}

func (m *spireMiddleware) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil || len(r.TLS.VerifiedChains) == 0 {
			http.Error(w, "spire: no verified client certificate", http.StatusUnauthorized)
			return
		}

		leaf := r.TLS.VerifiedChains[0][0]
		spiffeID, err := extractSPIFFEID(leaf)
		if err != nil {
			http.Error(w, "spire: invalid SVID", http.StatusUnauthorized)
			return
		}

		if len(m.allowed) > 0 {
			if _, ok := m.allowed[spiffeID]; !ok {
				logging.L.App.Warn("spire: SPIFFE ID not in allowlist", zap.String("id", spiffeID))
				http.Error(w, "spire: forbidden", http.StatusForbidden)
				return
			}
		}

		id := Identity{
			Method:   AuthMethodSPIRE,
			Subject:  leaf.Subject.String(),
			SPIFFEID: spiffeID,
		}

		next.ServeHTTP(w, WithIdentity(r, id))
	})
}

func extractSPIFFEID(cert *x509.Certificate) (string, error) {
	for _, uri := range cert.URIs {
		if uri.Scheme == "spiffe" {
			return uri.String(), nil
		}
	}
	return "", fmt.Errorf("no SPIFFE URI SAN found")
}
