package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type aesEncryptor struct {
	secretKey []byte
}

func NewAesEncryptor(keygen KeyGenerator, password string) aesEncryptor {
	return aesEncryptor{
		secretKey: keygen.Generate([]byte(password)),
	}
}

func (encryptor aesEncryptor) Encrypt(sensitiveData []byte) []byte {
	aes, err := aes.NewCipher(encryptor.secretKey)
	if err != nil {
		log.Fatalf("Error while creating new cipher for encrypt: %v", err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatalf("Error while creating new GCM for encrypt: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("Error while reading nonce: %v", err)
	}

	cipherText := gcm.Seal(nonce, nonce, sensitiveData, nil)

	return cipherText
}

func (encryptor aesEncryptor) Decrypt(encryptedData []byte) ([]byte, error) {
	aes, err := aes.NewCipher(encryptor.secretKey)
	if err != nil {
		log.Fatalf("Error while creating new cipher for decrypt: %v", err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatalf("Error while creating new GCM for decrypt: %v", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := encryptedData[:nonceSize], encryptedData[nonceSize:]

	decryptedData, err := gcm.Open(nil, nonce, cipherText, nil)

	return decryptedData, err
}
