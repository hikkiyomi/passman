package databases

import (
	"database/sql"
	"log"
	"os"

	"github.com/hikkiyomi/passman/internal/encryption"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	database  *sql.DB
	path      string
	encryptor encryption.Encryptor
}

// Initializes a database and creates table `storage` if it does not exist.
func newDatabase(database *sql.DB, path string, encryptor encryption.Encryptor) *Database {
	_, err := database.Exec(
		`
		CREATE TABLE IF NOT EXISTS storage (
			owner TEXT,
			service TEXT,
			data TEXT,
			PRIMARY KEY (owner, service)
		);
		`,
	)

	if err != nil {
		log.Fatalf("CREATE TABLE: %v", err)
	}

	return &Database{
		database:  database,
		path:      path,
		encryptor: encryptor,
	}
}

// Creates a connection to database.
// `path` is a path to your database, e.g. $HOME/passman/supersecret.db
// If you want to create a new database, put your desirable path to it as an argument.
func Open(path string, encryptor encryption.Encryptor) *Database {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Fatalf("Open error: %v", err)
	}

	return newDatabase(db, path, encryptor)
}

// Deletes the file where the entire database resides.
func (d *Database) Drop() {
	err := os.Remove(d.path)
	if err != nil {
		log.Fatalf("Drop error: %v", err)
	}
}

// Inserts a record into `storage` table.
func (d *Database) Insert(record Record) {
	encryptedRecord := record.Encrypt(d.encryptor)

	_, err := d.database.Exec(
		`
		INSERT INTO storage (owner, service, data)
		VALUES (?, ?, ?);
		`,
		encryptedRecord.Owner,
		encryptedRecord.Service,
		string(encryptedRecord.Data),
	)

	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}
}

// Returns all records from table `storage`.
func (d *Database) FindAll() []Record {
	rows, err := d.database.Query(
		`
		SELECT *
		FROM storage;
		`,
	)

	if err != nil {
		log.Fatalf("FindAll error while selecting: %v", err)
	}

	result := make([]Record, 0)

	for rows.Next() {
		var record Record

		if err := rows.Scan(&record.Owner, &record.Service, &record.Data); err != nil {
			log.Fatalf("Error while scanning with FindAll: %v", err)
		}

		record, err = record.Decrypt(d.encryptor)
		if err == nil {
			result = append(result, record)
		}
	}

	return result
}

// Retrieves records from `storage` table matching by owner.
func (d *Database) FindByOwner(owner string) []Record {
	rows, err := d.database.Query(
		`
		SELECT
			service,
			data
		FROM storage
		WHERE owner = ?;
		`,
		owner,
	)

	if err != nil {
		log.Fatalf("FindByOwner error while selecting: %v", err)
	}

	result := make([]Record, 0)

	for rows.Next() {
		record := Record{Owner: owner}

		if err := rows.Scan(&record.Service, &record.Data); err != nil {
			log.Fatalf("FindByOwner scanning error: %v", err)
		}

		record, err = record.Decrypt(d.encryptor)
		if err == nil {
			result = append(result, record)
		}
	}

	return result
}

// Retrieves the only tuple from `storage` table matching by owner and service,
// or nil if such does not exist.
func (d *Database) FindByOwnerAndService(owner, service string) *Record {
	var result *Record
	records := d.FindByOwner(owner)

	for _, record := range records {
		if record.Service == service {
			result = &record
			break
		}
	}

	return result
}

// This check is mandatory.
// Without it everybody can manipulate with your secure data.
// If user has provided incorrect master password or salt,
// they will be given nothing.
func (d *Database) checkIfRecordExists(owner, service string) bool {
	return d.FindByOwnerAndService(owner, service) != nil
}

// Updates the data in `storage` table matching by owner and service.
func (d *Database) Update(record Record) {
	if !d.checkIfRecordExists(record.Owner, record.Service) {
		return
	}

	encryptedRecord := record.Encrypt(d.encryptor)

	_, err := d.database.Exec(
		`
		UPDATE storage
		SET data = ?
		WHERE owner = ?
			AND service = ?;
		`,
		string(encryptedRecord.Data),
		encryptedRecord.Owner,
		encryptedRecord.Service,
	)

	if err != nil {
		log.Fatalf("Update error: %v", err)
	}
}

// Deletes all passwords from `storage` table matching by owner and service.
func (d *Database) DeleteByOwnerAndService(owner, service string) {
	if !d.checkIfRecordExists(owner, service) {
		return
	}

	_, err := d.database.Exec(
		`
		DELETE FROM storage
		WHERE owner = ?
			AND service = ?;
		`,
		owner,
		service,
	)

	if err != nil {
		log.Fatalf("Delete error: %v", err)
	}
}

// Deletes all passwords from `storage` table matching by owner and service.
// This method completely ignores data of the record.
func (d *Database) Delete(record Record) {
	d.DeleteByOwnerAndService(record.Owner, record.Service)
}
