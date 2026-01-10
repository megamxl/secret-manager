package types

type CreateSecretRequest struct {
	FilePath      string `json:"file_path,omitempty"`
	SecretWrapper Secret `json:"secret_wrapper"`
	StoreName     string `json:"store_name"`
}

type Secret struct {
	Content          string            `json:"content,omitempty"`
	SecretReferences map[string]string `json:"secret_references,omitempty"`
}
