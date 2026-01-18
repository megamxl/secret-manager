package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"secret-manager/pkg/service"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/types"
	"time"

	"github.com/onatm/clockwerk"
	"gorm.io/gorm"
	"sigs.k8s.io/yaml"
)

var tempService service.TemplateService
var storeService service.StoreService

func secretCreationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, ok := decodeSecretRequest(w, r)
	if !ok {
		return
	}

	// 1. Try to Fetch/Template FIRST (Validation)
	if handleServiceError(w, tempService.FetchAndStoreTemplate(req)) {
		return
	}

	// 2. Only if Fetch succeeded, we persist to DB
	if handleServiceError(w, tempService.StoreSecretConfig(req)) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func secretUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, ok := decodeSecretRequest(w, r)
	if !ok {
		return
	}

	if err := tempService.UpdateSecretConfig(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func secretDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name parameter is required", http.StatusBadRequest)
		return
	}

	if err := tempService.DeleteSecretConfig(name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleStoreCreate(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeStoreRequest(w, r)
	if !ok {
		return
	}

	if handleServiceError(w, storeService.CreateStore(req)) {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func handleStoreUpdate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Bad Request: name missing", http.StatusBadRequest)
		return
	}

	req, ok := decodeStoreRequest(w, r)
	if !ok {
		return
	}

	if handleServiceError(w, storeService.UpdateStore(name, req)) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleStoreDelete(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Bad Request: name missing", http.StatusBadRequest)
		return
	}

	if handleServiceError(w, storeService.DeleteStore(name)) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeStoreRequest(w http.ResponseWriter, r *http.Request) (stores.Config, bool) {
	var req stores.Config
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding store request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return req, false
	}
	defer r.Body.Close()
	return req, true
}

func decodeSecretRequest(w http.ResponseWriter, r *http.Request) (types.CreateSecretRequest, bool) {
	var req types.CreateSecretRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", 400)
		return req, false
	}
	defer r.Body.Close()

	if err := yaml.Unmarshal(body, &req); err != nil {
		log.Printf("Decode error: %v", err)
		http.Error(w, "Bad Request: Invalid format", 400)
		return req, false
	}

	return req, true
}

func handleServiceError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Printf("Service Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		secretCreationHandler(w, r)
	case http.MethodPut:
		secretUpdateHandler(w, r)
	case http.MethodDelete:
		secretDeleteHandler(w, r)
	default:
		w.Header().Set("Allow", "POST, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleStoreCreate(w, r)
	case http.MethodPut:
		handleStoreUpdate(w, r)
	case http.MethodDelete:
		handleStoreDelete(w, r)
	default:
		w.Header().Set("Allow", "POST, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SetupHandler(db *gorm.DB) {

	tempService = service.TemplateService{
		Db: db,
	}
	storeService = service.StoreService{
		Db: db,
	}

	job := service.RotationJob{
		Db:      db,
		Service: tempService,
	}

	c := clockwerk.New()
	c.Every(5 * time.Second).Do(job)
	c.Start()

}
