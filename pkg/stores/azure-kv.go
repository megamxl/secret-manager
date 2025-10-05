package stores

type az_kv struct {
}

var _ SecretBackend = (*az_kv)(nil)

func (a az_kv) instateClient() {

}

func (a az_kv) getSecrets() map[string]string {
	//TODO implement me
	panic("implement me")
}
