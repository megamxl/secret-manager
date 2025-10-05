package stores

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func SetupVaultClient(config *Config) (*vault.Client, error) {

	if config.Type != "hc-vault" {
		return nil, errors.New(fmt.Sprintf("Wrong type: %s only hc-Vaults", config.Type))
	}

	client, err := vault.New(
		vault.WithAddress(config.URL),
		vault.WithRequestTimeout(30*time.Second),
	)

	if err != nil {
		return nil, err
	}

	// authenticate with a root token (insecure)
	if config.Auth["token"] != "" {
		if err := client.SetToken(config.Auth["token"]); err != nil {
			//TODO change
			return nil, err
		}
		return client, nil
	}

	return nil, errors.New("AuthMethod not supported")
}

func GetSecretsFromDataStructure(pSecret ParsedSecret, mountPath string, client *vault.Client) (ResolvedSecretsWithMetaData, error) {

	type pathItem struct {
		before string
		after  string
	}

	// Collect all unique paths
	pathGroups := make(map[string][]pathItem)

	for _, ref := range pSecret.Refs {
		pathGroups[ref.PathToFolder] = append(pathGroups[ref.PathToFolder], pathItem{ref.PathToFolder, ref.RemoteSecretName})
	}

	result := make(map[string]string)
	appliedPaths := make(map[string]string)

	// Fetch secrets per path group
	for path := range pathGroups {
		secret, err := client.Secrets.KvV2Read(context.Background(), path, vault.WithMountPath(mountPath))
		if err != nil {
			log.Printf("Error reading path %s: %v\n", path, err)
			return ResolvedSecretsWithMetaData{}, err
		}

		data, err := toStringMap(secret.Data.Data)
		if err != nil {
			return ResolvedSecretsWithMetaData{}, err
		}

		for remoteName, secValue := range data {
			name := getRefByRemoteName(pSecret.Refs, path, remoteName)
			if name != nil {
				result[name.Name] = secValue
			}
		}

		if versionVal, ok := secret.Data.Metadata["version"]; ok {
			appliedPaths[path] = fmt.Sprintf("%v", versionVal)
		}
	}

	return ResolvedSecretsWithMetaData{
		Secrets: result,
		Paths:   appliedPaths,
	}, nil
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

func getRefByRemoteName(refs []SecretRefs, path string, remoteName string) *SecretRefs {

	for _, ref := range refs {
		if ref.getFullPath() == path+"/"+remoteName {
			return &ref
		}
	}

	return nil
}
