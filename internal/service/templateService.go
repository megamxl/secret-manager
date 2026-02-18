package service

import (
	"fmt"
	"log"
	"secret-manager/internal/logging"
	"secret-manager/pkg/files_helper"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/templating"
	"secret-manager/pkg/types"

	"gorm.io/gorm"
)

type SecretService interface {
	FetchAndStoreTemplate(req types.CreateSecretRequest) error
	StoreSecretConfig(req types.CreateSecretRequest) error
	UpdateSecretConfig(req types.CreateSecretRequest) error
	DeleteSecretConfig(name string) error
}
type TemplateServiceImpl struct {
	Db  *gorm.DB
	Log *logging.Loggers
}

func (t TemplateServiceImpl) StoreSecretConfig(req types.CreateSecretRequest) error {
	return persistence.SaveSecretConfig(t.Db, req)
}

func (t TemplateServiceImpl) UpdateSecretConfig(req types.CreateSecretRequest) error {
	err := persistence.UpdateSecretConfig(t.Db, req)
	if err != nil {
		return err
	}

	return t.FetchAndStoreTemplate(req)
}

func (t TemplateServiceImpl) DeleteSecretConfig(name string) error {
	// TODO: add logic here to delete the physical file

	byName, err := persistence.GetRequestByName(t.Db, name)
	if err != nil {
		return err
	}

	err = files_helper.DeleteFile(byName.FilePath)
	if err != nil {
		return err
	}

	return persistence.DeleteSecretConfig(t.Db, name)
}

func (t TemplateServiceImpl) FetchAndStoreTemplate(req types.CreateSecretRequest) error {

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

	template, err := templating.EvaluateTemplate(req.SecretWrapper.Content, resolvedSecrets, req.Name)
	if err != nil {
		log.Print("Error evaluating template:", err)
		return err
	}

	err = files_helper.CreateFileAtomic(req.FilePath, template)
	if err != nil {
		log.Print("Error creating file:", err)
		return err
	}

	t.Log.AuditLogUserEvent(fmt.Sprintf("File at path %s was resolved and saved", req.FilePath), "system", "rerolledFile")

	return nil
}

func ToParsedSecret(request types.CreateSecretRequest, config stores.Config) stores.ParsedSecret {

	refs := generateSecretRefsFromCreationRequest(request)

	return stores.ParsedSecret{
		FilePath:  request.FilePath,
		Content:   request.SecretWrapper.Content,
		StoreName: request.StoreName,
		MountPath: config.MountPath,
		Refs:      refs,
	}
}

func generateSecretRefsFromCreationRequest(request types.CreateSecretRequest) []stores.SecretRef {

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
