package service

import (
	"encoding/json"
	"io"
	"log"
	"secret-manager/pkg/files_helper"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/templating"

	"gorm.io/gorm"
)

type TemplateService struct {
	Db *gorm.DB
}

type CreateSecretRequest struct {
	FilePath      string `json:"file_path,omitempty"`
	SecretWrapper Secret `json:"secret_wrapper"`
	StoreName     string `json:"store_name"`
}

type Secret struct {
	Content          string            `json:"content,omitempty"`
	SecretReferences map[string]string `json:"secret_references,omitempty"`
}

func (t TemplateService) FetchAndStoreTemplate(r io.ReadCloser) error {

	var req CreateSecretRequest

	err := json.NewDecoder(r).Decode(&req)
	if err != nil {
		log.Print("Error decoding request:", err)
		return err
	}

	configByName, err := persistence.FindConfigByName(t.Db, req.StoreName)
	if err != nil {
		log.Print("Error loading secret-config.json:", err)
		return err
	}

	store, err := configByName.CreateBackend()
	if err != nil {
		log.Print("Error creating store:", err)
		return err
	}

	secret := ToParsedSecret(req, configByName)

	resolvedSecrets, err := store.ResolveSecrets(secret)
	if err != nil {
		log.Print("Error fetching secrets:", err)
		return err
	}

	template, err := templating.EvaluateTemplate(req.SecretWrapper.Content, resolvedSecrets)
	if err != nil {
		log.Print("Error evaluating template:", err)
		return err
	}

	err = files_helper.CreateFile(req.FilePath, template)
	if err != nil {
		log.Print("Error creating file:", err)
		return err
	}

	return nil
}

func ToParsedSecret(request CreateSecretRequest, config stores.Config) stores.ParsedSecret {

	refs := generateSecretRefsFromCreationRequest(request)

	return stores.ParsedSecret{
		FilePath:  request.FilePath,
		Content:   request.SecretWrapper.Content,
		StoreName: request.StoreName,
		MountPath: config.MountPath,
		Refs:      refs,
	}
}

func generateSecretRefsFromCreationRequest(request CreateSecretRequest) []stores.SecretRef {

	refs := make([]stores.SecretRef, 0)

	for k, v := range request.SecretWrapper.SecretReferences {

		path, remoteName := stores.SplitLastSlash(v)

		refs = append(refs, stores.SecretRef{
			TemplateName:                k,
			RemotePathWithoutSecretName: path,
			RemoteSecretName:            remoteName,
		})
	}

	return refs
}
