package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"secret-manager/pkg/service"
	"secret-manager/pkg/stores"

	"gorm.io/gorm"
)

var tempService service.TemplateService
var storeService service.StoreService

func SecretCreationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	err := tempService.FetchAndStoreTemplate(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	w.WriteHeader(http.StatusCreated)
}

func StoreCreationHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}

	var req stores.Config
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = storeService.CreateStore(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SetupHandler(db *gorm.DB) {

	tempService = service.TemplateService{
		Db: db,
	}
	storeService = service.StoreService{
		Db: db,
	}

}
