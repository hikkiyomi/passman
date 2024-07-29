package databases

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/hikkiyomi/passman/internal/encryption"
)

type Record struct {
	Owner   string
	Service string
	Data    []byte
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

func (r Record) Decrypt(encryptor encryption.Encryptor) (Record, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(r.Data))
	if err != nil {
		log.Fatal("Couldn't decode base64 data.")
	}

	resultingData, err := encryptor.Decrypt(decodedData)

	return Record{
		Owner:   r.Owner,
		Service: r.Service,
		Data:    resultingData,
	}, err
}

func (r Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Owner   string `json:"owner"`
		Service string `json:"service"`
		Data    string `json:"data"`
	}{
		Owner:   r.Owner,
		Service: r.Service,
		Data:    string(r.Data),
	})
}

func (r *Record) UnmarshalJSON(data []byte) error {
	var input *struct {
		Owner   string `json:"owner"`
		Service string `json:"service"`
		Data    string `json:"data"`
	}

	err := json.Unmarshal(data, &input)

	r.Owner = input.Owner
	r.Service = input.Service
	r.Data = []byte(input.Data)

	return err
}
