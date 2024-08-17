package databases_test

import (
	"log"
	"testing"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
)

var (
	kdf           = encryption.NewArgon2Kdf([]byte("salt"), 0, 0, 0, 0)
	aesEncryptor  = encryption.NewAesEncryptor(kdf, "password")
	noOpEncryptor = encryption.NoOpEncryptor{}
	databasePath  = "test.db"
	user          = "user"
)

// This test fails and that is expected.
// Running it will cause FAIL, because it will try to unmarshal some garbage.
func TestEncryption(t *testing.T) {
	localDatabasePath := "shouldNotBeDeleted.db"

	databaseWithEncryption := databases.Open(user, localDatabasePath, aesEncryptor)
	defer databaseWithEncryption.Drop()

	record := databases.Record{
		Owner:   user,
		Service: "some service",
		Data:    []byte("sike u thought"),
	}
	databaseWithEncryption.Insert(&record)

	databaseWithNoEncryption := databases.Open(user, localDatabasePath, &noOpEncryptor)

	databaseWithNoEncryption.FindByService("some service")
}

func TestInsert(t *testing.T) {
	collectionToInsert := []databases.Record{
		{
			Owner:   user,
			Service: "hehe",
			Data:    []byte("some data"),
		},
		{
			Owner:   user,
			Service: "another service",
			Data:    []byte("some data again"),
		},
		{
			Owner:   "you",
			Service: "your service",
			Data:    []byte("your data"),
		},
	}

	database := databases.Open(user, databasePath, aesEncryptor)
	defer database.Drop()

	for _, record := range collectionToInsert {
		database.Insert(&record)
	}

	if found := database.FindAll(); len(found) != 2 {
		t.Fatalf("expected 2 records with owner `me` but found %d", len(found))
	}

	if found := database.FindByService("another service"); len(found) == 0 {
		t.Fatal("expected finding the record with owner `me` and service `another service` but found nothing.")
	}

	anotherDatabase := databases.Open("you", databasePath, aesEncryptor)

	if found := anotherDatabase.FindAll(); len(found) != 1 {
		t.Fatalf("expected 1 record but found %v", len(found))
	}
}

func TestUpdate(t *testing.T) {
	database := databases.Open(user, databasePath, aesEncryptor)
	defer database.Drop()

	record := databases.Record{
		Owner:   user,
		Service: "service",
		Data:    []byte("kek"),
	}
	database.Insert(&record)

	record.Data = []byte("new kek")
	database.Update(record)

	if found := database.FindByService("service")[0]; string(found.Data) != string(record.Data) {
		t.Fatalf("expected %v but found %v", found.Data, record.Data)
	}
}

func TestDelete(t *testing.T) {
	database := databases.Open(user, databasePath, aesEncryptor)
	defer database.Drop()

	record := databases.Record{
		Owner:   user,
		Service: "service",
		Data:    []byte("data"),
	}

	database.Insert(&record)
	database.Delete(record)

	if found := database.FindAll(); len(found) != 0 {
		log.Fatalf("expected 0 records in database but found %d", len(found))
	}
}
