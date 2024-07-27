package encryption

// Do not use that. NoOpEncryptor was created specifically for tests.
type NoOpEncryptor struct {
}

func (encryptor *NoOpEncryptor) Encrypt(data []byte) []byte {
	return data
}

func (encryptor *NoOpEncryptor) Decrypt(data []byte) []byte {
	return data
}
