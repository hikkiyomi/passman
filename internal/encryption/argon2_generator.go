package encryption

import "golang.org/x/crypto/argon2"

type Argon2Kdf struct {
	salt    []byte
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func NewArgon2Kdf(
	salt []byte,
	time uint32,
	memory uint32,
	threads uint8,
	keyLen uint32,
) Argon2Kdf {
	return Argon2Kdf{salt, time, memory, threads, keyLen}
}

func (kdf Argon2Kdf) Generate(data []byte) []byte {
	return argon2.Key(data, kdf.salt, kdf.time, kdf.memory, kdf.threads, kdf.keyLen)
}
