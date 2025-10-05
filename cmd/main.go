package main

import (
	"log"
	"net/http"
	"secret-manager/pkg/handlers"
	"secret-manager/pkg/persistence"
)

func main() {

	var err error

	db, err := persistence.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	handlers.SetupHandler(db)

	http.HandleFunc("/create", handlers.SecretCreationHandler)
	http.HandleFunc("/store", handlers.StoreCreationHandler)

	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}
