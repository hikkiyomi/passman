package internal_test

import (
	"log"
	"testing"

	"github.com/hikkiyomi/passman/internal"
	"github.com/hikkiyomi/passman/internal/encryption"
)

var (
	kdf           = encryption.NewArgon2Kdf([]byte("salt"), 0, 0, 0, 0)
	aesEncryptor  = encryption.NewAesEncryptor(kdf, "password")
	noOpEncryptor = encryption.NoOpEncryptor{}
)

func TestEncryption(t *testing.T) {
	databaseWithEncryption := internal.Open(".", aesEncryptor)
	defer databaseWithEncryption.Drop()

	record := internal.Record{
		Owner:   "me",
		Service: "some service",
		Data:    []byte("sike u thought"),
	}
	databaseWithEncryption.Insert(record)

	databaseWithNoEncryption := internal.Open(".", &noOpEncryptor)
	foundRecord := databaseWithNoEncryption.FindByOwnerAndService("me", "some service")

	if foundRecord.Owner != record.Owner ||
		foundRecord.Service != record.Service ||
		string(foundRecord.Data) == string(record.Data) {

		t.Fatalf("expected encrypted data %s. found the same %s.", foundRecord.Data, record.Data)
	}

	foundRecordDecrypted := databaseWithEncryption.FindByOwnerAndService("me", "some service")

	if foundRecordDecrypted.Owner != record.Owner ||
		foundRecordDecrypted.Service != record.Service ||
		string(foundRecordDecrypted.Data) != string(record.Data) {

		t.Fatalf("expected %v, but found %v", record, foundRecordDecrypted)
	}
}

func TestInsert(t *testing.T) {
	collectionToInsert := []internal.Record{
		{
			Owner:   "me",
			Service: "hehe",
			Data:    []byte("some data"),
		},
		{
			Owner:   "me",
			Service: "another service",
			Data:    []byte("some data again"),
		},
		{
			Owner:   "you",
			Service: "your service",
			Data:    []byte("your data"),
		},
	}

	database := internal.Open(".", aesEncryptor)
	defer database.Drop()

	for _, record := range collectionToInsert {
		database.Insert(record)
	}

	if found := database.FindByOwner("me"); len(found) != 2 {
		t.Fatalf("expected 2 records with owner `me` but found %d", len(found))
	}

	if found := database.FindByOwnerAndService("me", "another service"); found == nil {
		t.Fatal("expected finding the record with owner `me` and service `another service` but found nothing.")
	}

	if found := database.FindByOwner("you"); len(found) != 1 {
		t.Fatalf("expected 2 records with owner `you` but found %d", len(found))
	}

	if found := database.FindAll(); len(found) != 3 {
		t.Fatalf("expected 3 records in total but found %d", len(found))
	}
}

func TestUpdate(t *testing.T) {
	database := internal.Open(".", aesEncryptor)
	defer database.Drop()

	record := internal.Record{
		Owner:   "me",
		Service: "service",
		Data:    []byte("kek"),
	}
	database.Insert(record)

	record.Data = []byte("new kek")
	database.Update(record)

	if found := database.FindByOwnerAndService("me", "service"); string(found.Data) != string(record.Data) {
		t.Fatalf("expected %v but found %v", found.Data, record.Data)
	}
}

func TestDelete(t *testing.T) {
	database := internal.Open(".", aesEncryptor)
	defer database.Drop()

	record := internal.Record{
		Owner:   "me",
		Service: "service",
		Data:    []byte("data"),
	}

	database.Insert(record)
	database.Delete(record)

	if found := database.FindAll(); len(found) != 0 {
		log.Fatalf("expected 0 records in database but found %d", len(found))
	}
}
