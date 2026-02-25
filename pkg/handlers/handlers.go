package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := tempService.UpdateSecretConfig(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func secretDeleteHandler(w http.ResponseWriter, r *http.Request, req types.CreateSecretRequest) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := tempService.DeleteSecretConfig(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

func decodeStoreRequest(w http.ResponseWriter, r *http.Request) (stores.Config, bool) {
	var req stores.Config
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logging.L.App.Error(fmt.Sprintf("Error decoding store request: %v", err))
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
		logging.L.App.Error(fmt.Sprintf("Error decoding request: %v", err))
		http.Error(w, "Bad Request: Invalid format", 400)
		return req, false
	}

	return req, true
}

func handleServiceError(w http.ResponseWriter, err error) bool {
	if err != nil {
		logging.L.App.Error(fmt.Sprintf("Service Error: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {

	req, ok := decodeSecretRequest(w, r)
	if !ok && r.Method != http.MethodDelete {
		http.Error(w, "Bad Request Not valid SecretRequest", http.StatusBadRequest)
		return
	}

	action := ""

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
	default:
		w.Header().Set("Allow", "POST, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	caller, err := callerString(r)
	if err != nil {
		http.Error(w, "The request passed the Auth Chack but no Identify was set call the Support", http.StatusBadRequest)
	}

	logging.L.AuditLogUserEvent(
		fmt.Sprintf("Secret with the name %s was changed", req.Name),
		caller,
		action,
	)
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {

	req, ok := decodeStoreRequest(w, r)
	if !ok && r.Method != http.MethodDelete {
		http.Error(w, "Bad Request Not valid SecretRequest", http.StatusBadRequest)
		return
	}

	action := ""

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
	default:
		w.Header().Set("Allow", "POST, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	caller, err := callerString(r)
	if err != nil {
		http.Error(w, "The request passed the Auth Chack but no Identify was set call the Support", http.StatusBadRequest)
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
