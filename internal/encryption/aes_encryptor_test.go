package encryption_test

import (
	"testing"

	"github.com/hikkiyomi/passman/internal/encryption"
)

var (
	kdf                       = encryption.NewArgon2Kdf([]byte("salt"), 0, 0, 0, 0)
	encryptor                 = encryption.NewAesEncryptor(kdf, "password")
	collectionOfSensitiveData = [][]byte{
		[]byte("some sensitive data lol"),
		[]byte("bangboos are cute"),
		[]byte("who dis"),
		[]byte("totally not a password"),
		[]byte("somebody once told me"),
	}
)

func TestEncrypt(t *testing.T) {
	for _, sensitiveData := range collectionOfSensitiveData {
		encryptedData := encryptor.Encrypt(sensitiveData)

		if string(sensitiveData) == string(encryptedData) {
			t.Fatalf("data `%s` was not encrypted", sensitiveData)
		}
	}
}

func TestDecrypt(t *testing.T) {
	for _, sensitiveData := range collectionOfSensitiveData {
		encryptedData := encryptor.Encrypt(sensitiveData)
		decryptedData, _ := encryptor.Decrypt(encryptedData)

		if string(sensitiveData) != string(decryptedData) {
			t.Fatalf("decrypted data `%s` is not equal to original data `%s`", decryptedData, sensitiveData)
		}
	}
}
