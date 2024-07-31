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
	id      int64
}

type EncryptedRecord struct {
	Data []byte
	id   int64
}

func (r Record) Encrypt(encryptor encryption.Encryptor) EncryptedRecord {
	fullData, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("Error while marshalling record for encrypting: %v", err)
	}

	encryptedData := encryptor.Encrypt(fullData)
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	return EncryptedRecord{
		id:   r.id,
		Data: []byte(encodedData),
	}
}

func (r EncryptedRecord) Decrypt(encryptor encryption.Encryptor) (Record, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(r.Data))
	if err != nil {
		log.Fatalf("Couldn't decode base64 data: %v", err)
	}

	resultingData, err := encryptor.Decrypt(decodedData)
	var result Record

	if err == nil {
		jsonErr := json.Unmarshal(resultingData, &result)
		if jsonErr != nil {
			log.Fatalf("Couldn't unmarshal json: %v", jsonErr)
		}

		result.id = r.id
	}

	return result, err
}

func (r Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id      int64  `json:"id"`
		Owner   string `json:"owner"`
		Service string `json:"service"`
		Data    string `json:"data"`
	}{
		Id:      r.id,
		Owner:   r.Owner,
		Service: r.Service,
		Data:    string(r.Data),
	})
}

func (r *Record) UnmarshalJSON(data []byte) error {
	var input *struct {
		Id      int64  `json:"id"`
		Owner   string `json:"owner"`
		Service string `json:"service"`
		Data    string `json:"data"`
	}

	err := json.Unmarshal(data, &input)

	r.id = input.Id
	r.Owner = input.Owner
	r.Service = input.Service
	r.Data = []byte(input.Data)

	return err
}
