package stores

type ResolvedSecretsWithMetaData struct {
	Secrets map[string]string
	Paths   map[string]string
}

type SecretBackend interface {
	instateClient()
	getSecrets() map[string]string
}

type ParsedSecret struct {
	FilePath  string       `json:"file_path,omitempty"`
	Content   string       `json:"content,omitempty"`
	StoreName string       `json:"store_name"`
	Refs      []SecretRefs `json:"refs"`
}

type SecretRefs struct {
	Name             string `json:"secret_name"`
	PathToFolder     string `json:"secret_path_to_folder"`
	RemoteSecretName string `json:"secret_remote_secret_name"`
}

func (ref SecretRefs) getFullPath() string {
	return ref.PathToFolder + "/" + ref.RemoteSecretName
}
