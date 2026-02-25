package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"secret-manager/internal/config"
	"secret-manager/internal/logging"
	"secret-manager/internal/service"
	"secret-manager/pkg/authentication"
	"secret-manager/pkg/handlers"
	"secret-manager/pkg/persistence"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {

	setupConfig()

	setupLoggers(config.Get())
	defer logging.L.Close()

	db, err := persistence.InitDB()
	if err != nil {
		logging.L.App.Fatal("db init failed", zap.Error(err))
	}

	handlers.SetupHandler(db)

	srv, err := BuildServer(config.Get())
	if err != nil {
		logging.L.App.Fatal("failed to initialize server", zap.Error(err))
	}

	runServer(srv)

}

func runServer(srv *http.Server) {
	logging.L.App.Info("server starting", zap.String("addr", srv.Addr))

	var err error
	if srv.TLSConfig != nil {
		// Certificates are already loaded in TLSConfig
		err = srv.ListenAndServeTLS("", "")
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logging.L.App.Fatal("server crashed", zap.Error(err))
	}

}

func BuildServer(cfg *config.Config) (*http.Server, error) {
	// 1. Setup Auth
	authMW, err := buildAuthentication(cfg.SecurityConfig.Auth)
	if err != nil {
		return nil, err
	}

	logging.L.Audit.Info(fmt.Sprintf("Auth is set to %s", cfg.SecurityConfig.Auth.Method), zap.String("action", "SecurityConfig"))

	// 2. Setup TLS
	tlsCfg, err := buildTLSConfig(cfg.SecurityConfig)
	if err != nil {
		return nil, err
	}

	// 3. Setup Routing (Plain Handler)
	// We wrap the router in the auth authentication so it's global
	handler := authMW(http.HandlerFunc(routeRequest))

	setupPrometheus()

	return &http.Server{
		Addr:      fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:   handler,
		TLSConfig: tlsCfg,
	}, nil
}

func setupPrometheus() *prometheus.Registry {
	// Metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		service.RerollSelectedTotal,
		service.RerollFailureTotal,
		service.RerollSuccessTotal,
	)
	return reg
}

func setupLoggers(cfg *config.Config) {
	logging.Setup(cfg.Logging.Level, cfg.Logging.Format, cfg.Logging.AuditFile)

	logging.L.App.Info("starting secret-manager",
		zap.String("node", cfg.Agent.NodeName),
		zap.Int("port", cfg.Server.Port),
	)
}

func setupConfig() {
	configPath := flag.String("config", "config.yaml", "Config file path")
	flag.Parse()

	if err := config.Initialize(*configPath); err != nil {
		panic(err)
	}
}

// routeRequest acts as your manual router
func routeRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/metrics":
		// Handle metrics (Note: reg should be globally accessible or passed in)
		promhttp.Handler().ServeHTTP(w, r)
	case "/secret":
		handlers.SecretHandler(w, r)
	case "/store":
		handlers.StoreHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func buildAuthentication(cfg config.AuthConfig) (func(http.Handler) http.Handler, error) {
	switch cfg.Method {
	case "none", "":
		logging.L.App.Warn("Auth disabled running not secure only for Test cases !!!")
		logging.L.Audit.Warn("Auth is-disabled running not secure only for Test cases !!!", zap.String("action", "SecurityConfig"))
		return authentication.NoAuthMiddleware, nil
	case "mtls":
		return authentication.MTLSMiddleware, nil
	case "jwt":
		return authentication.NewJWTMiddleware(cfg.JWT)
	case "spire":
		return authentication.NewSPIREMiddleware(cfg.SPIRE)
	default:
		return nil, fmt.Errorf("unknown auth method: %q", cfg.Method)
	}
}

func buildTLSConfig(cfg config.SecurityConfig) (*tls.Config, error) {
	if !cfg.TLS.Enabled {
		return nil, nil
	}

	if cfg.Auth.Method == "spire" {
		return authentication.SPIRETLSConfig(
			cfg.Auth.SPIRE.SocketPath,
			cfg.Auth.SPIRE.TrustDomain,
		)
	}
	return authentication.BuildTLSConfig(cfg.TLS)
}
