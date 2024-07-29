package encryption

type Encryptor interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) ([]byte, error)
}
