package stores

type ResolvedSecretsWithMetaData struct {
	Secrets map[string]string
	Paths   map[string]string
}

type SecretBackend interface {
	InstateClient(config *Config) error
	ResolveSecrets(ParsedSecret) (map[string]string, error)
}

type ParsedSecret struct {
	FilePath  string      `json:"file_path,omitempty"`
	Content   string      `json:"content,omitempty"`
	StoreName string      `json:"store_name"`
	MountPath string      `json:"mount_path"`
	Refs      []SecretRef `json:"refs"`
}

type SecretRef struct {
	TemplateName                string `json:"secret_template_name"`
	RemotePathWithoutSecretName string `json:"secret_remote_path_without_secret_name"`
	RemoteSecretName            string `json:"secret_remote_secret_name"`
}
