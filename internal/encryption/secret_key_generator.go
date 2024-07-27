package encryption

type KeyGenerator interface {
	Generate([]byte) []byte
}
