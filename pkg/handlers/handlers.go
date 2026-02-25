package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"secret-manager/internal/helpers"
	"secret-manager/internal/logging"
	"secret-manager/internal/service"
	"secret-manager/pkg/authentication"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/types"
	"time"

	"github.com/onatm/clockwerk"
	"gorm.io/gorm"
	"sigs.k8s.io/yaml"
)

var tempService service.SecretService
var storeService service.StoreService

func secretCreationHandler(w http.ResponseWriter, r *http.Request, req types.CreateSecretRequest) {
	if r.Method != http.MethodPost {
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
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

func secretUpdateHandler(w http.ResponseWriter, r *http.Request, req types.CreateSecretRequest) {
	if r.Method != http.MethodPut {
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := tempService.UpdateSecretConfig(req); err != nil {
		helpers.HttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func secretDeleteHandler(w http.ResponseWriter, r *http.Request, req types.CreateSecretRequest) {
	if r.Method != http.MethodDelete {
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := tempService.DeleteSecretConfig(req.Name); err != nil {
		helpers.HttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func secretListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	secrets, err := tempService.GetAllSecretConfigs()
	if handleServiceError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(secrets); err != nil {
		logging.L.App.Error(fmt.Sprintf("Error encoding secrets: %v", err))
		helpers.HttpError(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func handleStoreCreate(w http.ResponseWriter, r *http.Request, req stores.Config) {

	if handleServiceError(w, storeService.CreateStore(req)) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleStoreUpdate(w http.ResponseWriter, r *http.Request, req stores.Config) {

	if handleServiceError(w, storeService.UpdateStore(req)) {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleStoreDelete(w http.ResponseWriter, r *http.Request, req stores.Config) {

	if handleServiceError(w, storeService.DeleteStore(req.ReferenceName)) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func storeListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	allStores, err := storeService.GetAllStores()
	if handleServiceError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allStores); err != nil {
		logging.L.App.Error(fmt.Sprintf("Error encoding stores: %v", err))
		helpers.HttpError(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func decodeStoreRequest(w http.ResponseWriter, r *http.Request) (stores.Config, bool) {
	var req stores.Config
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logging.L.App.Error(fmt.Sprintf("Error decoding store request: %v", err))
		helpers.HttpError(w, "Bad Request", http.StatusBadRequest)
		return req, false
	}
	defer r.Body.Close()
	return req, true
}

func decodeSecretRequest(w http.ResponseWriter, r *http.Request) (types.CreateSecretRequest, bool) {
	var req types.CreateSecretRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.HttpError(w, "Can't read body", 400)
		return req, false
	}
	defer r.Body.Close()

	if err := yaml.Unmarshal(body, &req); err != nil {
		logging.L.App.Error(fmt.Sprintf("Error decoding request: %v", err))
		helpers.HttpError(w, "Bad Request: Invalid format", 400)
		return req, false
	}

	return req, true
}

func handleServiceError(w http.ResponseWriter, err error) bool {
	if err != nil {
		logging.L.App.Error(fmt.Sprintf("Service Error: %v", err))
		helpers.HttpError(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {

	var req types.CreateSecretRequest
	var ok bool // Declare ok beforehand
	action := ""

	if r.Method != http.MethodGet {
		req, ok = decodeSecretRequest(w, r)
		if !ok {
			helpers.HttpError(w, "Bad Request Not valid SecretRequest", http.StatusBadRequest)
			return
		}
	}

	switch r.Method {
	case http.MethodPost:
		secretCreationHandler(w, r, req)
		action = "CreateSecret"
	case http.MethodPut:
		secretUpdateHandler(w, r, req)
		action = "UpdateSecret"
	case http.MethodDelete:
		secretDeleteHandler(w, r, req)
		action = "DeleteSecret"
	case http.MethodGet:
		secretListHandler(w, r)
		action = "ListSecrets"
	default:
		w.Header().Set("Allow", "POST, PUT, DELETE")
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	caller, err := callerString(r)
	if err != nil {
		helpers.HttpError(w, "The request passed the Auth Chack but no Identify was set call the Support", http.StatusBadRequest)
	}

	if w.Header().Get("IsError") == "true" {
		logging.L.AuditLogUserEvent(
			fmt.Sprintf("FAILED update secret %s", req.FilePath),
			caller,
			action)
		return
	}

	logging.L.AuditLogUserEvent(
		fmt.Sprintf("Secret with the name %s was changed", req.Name),
		caller,
		action,
	)
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {

	var req stores.Config
	var ok bool
	action := ""

	if r.Method != http.MethodGet {
		req, ok = decodeStoreRequest(w, r)
		if !ok {
			helpers.HttpError(w, "Bad Request Not valid SecretRequest", http.StatusBadRequest)
			return
		}
	}

	switch r.Method {
	case http.MethodPost:
		handleStoreCreate(w, r, req)
		action = "CreateStore"
	case http.MethodPut:
		handleStoreUpdate(w, r, req)
		action = "UpdateStore"
	case http.MethodDelete:
		handleStoreDelete(w, r, req)
		action = "DeleteStore"
	case http.MethodGet:
		storeListHandler(w, r)
		action = "ListStore"
	default:
		w.Header().Set("Allow", "POST, PUT, DELETE")
		helpers.HttpError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	caller, err := callerString(r)
	if err != nil {
		helpers.HttpError(w, "The request passed the Auth Chack but no Identify was set call the Support", http.StatusBadRequest)
	}

	if w.Header().Get("IsError") == "true" {
		logging.L.AuditLogUserEvent(
			fmt.Sprintf("FAILED update store %s", req.ReferenceName),
			caller,
			action)
		return
	}

	logging.L.AuditLogUserEvent(
		fmt.Sprintf("Secret with the name %s was changed", req.ReferenceName),
		caller,
		action,
	)
}

func callerString(r *http.Request) (string, error) {
	id, err := authentication.IdentityFrom(r.Context())
	if err != nil {
		return "", err
	}

	switch id.Method {
	case authentication.AuthMethodSPIRE:
		return fmt.Sprintf("spire:%s", id.SPIFFEID), nil
	case authentication.AuthMethodJWT:
		return fmt.Sprintf("jwt:%s", id.Subject), nil
	case authentication.AuthMethodMTLS:
		return fmt.Sprintf("mtls:%s", id.Subject), nil
	case authentication.AuthMethodNone:
		return fmt.Sprintf("ip:%s", r.RemoteAddr), nil
	default:
		return "", fmt.Errorf("unknown auth method: %q", id.Method)
	}
}

func SetupHandler(db *gorm.DB) {

	tempService = service.TemplateServiceImpl{
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
