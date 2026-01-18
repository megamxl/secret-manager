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

	http.HandleFunc("/secret", handlers.SecretHandler)
	http.HandleFunc("/store", handlers.StoreHandler)

	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}
