package encryption

import "golang.org/x/crypto/argon2"

type argon2Kdf struct {
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
) argon2Kdf {
	if time == 0 {
		time = 3
	}

	if memory == 0 {
		memory = 32 * 1024
	}

	if threads == 0 {
		threads = 4
	}

	if keyLen == 0 {
		keyLen = 32
	}

	return argon2Kdf{salt, time, memory, threads, keyLen}
}

func (kdf argon2Kdf) Generate(data []byte) []byte {
	return argon2.Key(data, kdf.salt, kdf.time, kdf.memory, kdf.threads, kdf.keyLen)
}
