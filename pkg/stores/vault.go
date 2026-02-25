package stores

import (
	"context"
	"errors"
	"fmt"
	"secret-manager/internal/logging"
	"strings"
	"time"

	"github.com/hashicorp/vault-client-go"
)

type VaultBackend struct {
	client *vault.Client
}

var _ SecretBackend = (*VaultBackend)(nil)

func (v *VaultBackend) InstateClient(config *Config) error {

	if config.Type != "hc-vault" {
		return errors.New(fmt.Sprintf("Wrong type: %s only hc-Vaults", config.Type))
	}

	client, err := vault.New(
		vault.WithAddress(config.URL),
		vault.WithRequestTimeout(30*time.Second),
	)

	if err != nil {
		return err
	}

	// authenticate with a root token (insecure)
	if config.Auth["token"] != "" {
		if err := client.SetToken(config.Auth["token"]); err != nil {
			return err
		}
		v.client = client
		return nil
	}

	return errors.New("AuthMethod not supported")
}

func (v *VaultBackend) ResolveSecrets(secret ParsedSecret) (map[string]string, error) {
	type pathItem struct {
		before string
		after  string
	}

	if v.client == nil {
		return nil, errors.New("vault client not initialized")
	}

	// Collect all unique paths
	pathGroups := make(map[string][]pathItem)

	for _, ref := range secret.Refs {
		pathGroups[ref.RemotePathWithoutSecretName] = append(pathGroups[ref.RemotePathWithoutSecretName], pathItem{ref.RemotePathWithoutSecretName, ref.RemoteSecretName})
	}

	result := make(map[string]string)
	appliedPaths := make(map[string]string)

	// Fetch secrets per path group
	for path := range pathGroups {
		secretResp, err := v.client.Secrets.KvV2Read(context.Background(), path, vault.WithMountPath(secret.MountPath))
		if err != nil {
			logging.L.App.Error(fmt.Sprintf("Error reading path %s: %v\n", path, err))
			return nil, err
		}

		data, err := toStringMap(secretResp.Data.Data)
		if err != nil {
			return nil, err
		}

		for remoteName, secValue := range data {
			// We iterate the refs to find which template key matches this specific path+key
			for _, ref := range secret.Refs {
				if ref.RemotePathWithoutSecretName == path && ref.RemoteSecretName == remoteName {
					result[ref.TemplateName] = secValue
				}
			}
		}

		if versionVal, ok := secretResp.Data.Metadata["version"]; ok {
			appliedPaths[path] = fmt.Sprintf("%v", versionVal)
		}
	}

	return result, nil
}

func toStringMap(m map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string, len(m))
	for k, v := range m {
		strVal, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("value for key %q is not a string (type %T)", k, v)
		}
		result[k] = strVal
	}
	return result, nil
}

func SplitLastSlash(s string) (before, after string) {
	lastSlash := strings.LastIndex(s, "/")
	if lastSlash == -1 {
		return "", s
	}
	return s[:lastSlash], s[lastSlash+1:]
}

func getRefByRemoteName(refs []SecretRef, path string, remoteName string) *SecretRef {

	for _, ref := range refs {
		if getFullPathForVault(ref) == path+"/"+remoteName {
			return &ref
		}
	}

	return nil
}

func getFullPathForVault(ref SecretRef) string {
	return ref.RemotePathWithoutSecretName + "/" + ref.RemoteSecretName
}
