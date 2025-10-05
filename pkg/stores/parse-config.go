package stores

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/vault-client-go"
)

type Config struct {
	Type          string            `json:"type" yaml:"type"`
	URL           string            `json:"url" yaml:"url"`
	Auth          map[string]string `json:"auth" yaml:"auth"`
	MountPath     string            `json:"mountPath" yaml:"mountPath"`
	ReferenceName string            `json:"referenceName" yaml:"referenceName"`
}

func LoadConfigFromFile(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)

	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

func LoadConfigFromEnv() {}

func (c Config) CreateStore() (*vault.Client, error) {

	switch c.Type {
	case "hc-vault":
		client, err := SetupVaultClient(&c)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	return nil, errors.New("invalid vault type")
}
