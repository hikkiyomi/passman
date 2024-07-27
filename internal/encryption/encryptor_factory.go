package encryption

func GetEncryptor(chosenEncryptor, masterPassword, salt string) Encryptor {
	var result Encryptor

	switch chosenEncryptor {
	case "aes":
		kdf := NewArgon2Kdf([]byte(salt), 0, 0, 0, 0)
		result = NewAesEncryptor(kdf, masterPassword)
	}

	return result
}
