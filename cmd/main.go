package main

import (
	"log"
	"net/http"
	"secret-manager/pkg/handlers"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/service"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	var err error

	db, err := persistence.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	handlers.SetupHandler(db)

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		service.RerollSelectedTotal,
		service.RerollFailureTotal,
		service.RerollSuccessTotal,
	)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	http.HandleFunc("/secret", handlers.SecretHandler)
	http.HandleFunc("/store", handlers.StoreHandler)

	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}
