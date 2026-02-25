package main

import (
	"flag"
	"fmt"
	"net/http"
	"secret-manager/internal/config"
	"secret-manager/internal/logging"
	"secret-manager/internal/service"
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

	reg := setupPrometheus()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.HandleFunc("/secret", handlers.SecretHandler)
	http.HandleFunc("/store", handlers.StoreHandler)

	addr := fmt.Sprintf(":%d", config.Get().Server.Port)
	logging.L.App.Info("server listening", zap.String("addr", addr))

	if err := http.ListenAndServe(addr, nil); err != nil {
		logging.L.App.Fatal("server failed", zap.Error(err))
	}

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
