package internal

import (
	"database/sql"
	"encoding/base64"
	"log"

	"github.com/hikkiyomi/passman/internal/encryption"
)

type Record struct {
	Owner   string
	Service string
	Data    []byte
}

func (r *Record) Scan(rows *sql.Rows, encryptor encryption.Encryptor, scanData ...any) error {
	err := rows.Scan(scanData...)

	r.Service = service
	r.Data = []byte(encryptedData)
	*r = r.Decrypt(encryptor)

	return err
}

func (r Record) Encrypt(encryptor encryption.Encryptor) Record {
	encryptedData := encryptor.Encrypt(r.Data)
	resultingData := base64.StdEncoding.EncodeToString(encryptedData)

	return Record{
		Owner:   r.Owner,
		Service: r.Service,
		Data:    []byte(resultingData),
	}
}

func (r Record) Decrypt(encryptor encryption.Encryptor) Record {
	decodedData, err := base64.StdEncoding.DecodeString(string(r.Data))
	if err != nil {
		log.Fatal(err)
	}

	resultingData := encryptor.Decrypt(decodedData)

	return Record{
		Owner:   r.Owner,
		Service: r.Service,
		Data:    resultingData,
	}
}
