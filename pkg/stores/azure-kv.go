package stores

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type AzureKeyVaultBackend struct {
	client *azsecrets.Client
}

func (a *AzureKeyVaultBackend) InstateClient(config *Config) error {

	if config.Auth["tenant_id"] == "" || config.Auth["client_id"] == "" || config.Auth["client_secret"] == "" {
		return errors.New("missing credentials, you need tenant_id, client_id, and client_secret ")
	}

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewClientSecretCredential(config.Auth["tenant_id"], config.Auth["client_id"], config.Auth["client_secret"], nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(config.URL, cred, nil)
	if err != nil {
		return err
	}
	a.client = client

	return nil
}

func (a *AzureKeyVaultBackend) ResolveSecrets(secret ParsedSecret) (map[string]string, error) {

	if a.client == nil {
		return nil, errors.New("azure Key Vault client is nil")
	}

	result := make(map[string]string)

	for _, ref := range secret.Refs {

		response, err := a.client.GetSecret(context.Background(), ref.RemoteSecretName, "", nil)
		if err != nil {
			return nil, err
		}

		result[ref.TemplateName] = *response.Value
	}

	return result, nil
}

var _ SecretBackend = (*AzureKeyVaultBackend)(nil)
