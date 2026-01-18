package service

import (
	"log"
	"secret-manager/pkg/files_helper"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/templating"
	"secret-manager/pkg/types"

	"gorm.io/gorm"
)

type TemplateService struct {
	Db *gorm.DB
}

func (t TemplateService) StoreSecretConfig(req types.CreateSecretRequest) error {
	return persistence.SaveSecretConfig(t.Db, req)
}

func (t TemplateService) UpdateSecretConfig(req types.CreateSecretRequest) error {
	err := persistence.UpdateSecretConfig(t.Db, req)
	if err != nil {
		return err
	}

	return t.FetchAndStoreTemplate(req)
}

func (t TemplateService) DeleteSecretConfig(name string) error {
	// TODO: add logic here to delete the physical file
	//files_helper.DeleteFile()

	return persistence.DeleteSecretConfig(t.Db, name)
}

func (t TemplateService) FetchAndStoreTemplate(req types.CreateSecretRequest) error {

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
