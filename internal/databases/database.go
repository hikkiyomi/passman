package databases

import (
	"database/sql"
	"log"
	"os"

	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/hikkiyomi/passman/internal/util"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	database  *sql.DB
	path      string
	encryptor encryption.Encryptor
	user      string
}

// Initializes a database and creates table `storage` if it does not exist.
func newDatabase(database *sql.DB, user, path string, encryptor encryption.Encryptor) *Database {
	_, err := database.Exec(
		`
		CREATE TABLE IF NOT EXISTS storage (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			data TEXT
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
		user:      user,
	}
}

// Creates a connection to database.
// `path` is a path to your database, e.g. $HOME/passman/supersecret.db
// If you want to create a new database, put your desirable path to it as an argument.
func Open(user, path string, encryptor encryption.Encryptor) *Database {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Fatalf("Open error: %v", err)
	}

	return newDatabase(db, user, path, encryptor)
}

// Deletes the file where the entire database resides.
func (d *Database) Drop() {
	err := os.Remove(d.path)
	if err != nil {
		log.Fatalf("Drop error: %v", err)
	}
}

// Inserts a record into `storage` table.
func (d *Database) Insert(record *Record) {
	encryptedRecord := record.Encrypt(d.encryptor)

	insertResult, err := d.database.Exec(
		`
		INSERT INTO storage (data)
		VALUES (?);
		`,
		string(encryptedRecord.Data),
	)

	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}

	record.Id, err = insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("Couldn't change last inserted id: %v", err)
	}
}

// Returns all records from table `storage` that belong to current user.
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
		var encryptedRecord EncryptedRecord

		if err := rows.Scan(&encryptedRecord.id, &encryptedRecord.Data); err != nil {
			log.Fatalf("Error while scanning with FindAll: %v", err)
		}

		record, err := encryptedRecord.Decrypt(d.encryptor)
		if err == nil && record.Owner == d.user {
			result = append(result, record)
		}
	}

	return result
}

// Retrieves records from `storage` table matching by service.
func (d *Database) FindByService(service string) []Record {
	return util.Filter(d.FindAll(), func(x Record) bool {
		return x.Service == service
	})
}

// Retrieves records from `storage` table matching by id.
func (d *Database) FindById(id int64) *Record {
	var result *Record

	for _, record := range d.FindAll() {
		if record.Id == id {
			result = &record
			break
		}
	}

	return result
}

// This function is mandatory.
// Without it everybody can manipulate your secure data.
// If user has provided incorrect master password or salt, they will be given nothing.
// Therefore, they wouldn't be able to update/delete anything.
func (d *Database) exists(id int64) bool {
	return d.FindById(id) != nil
}

// Updates the data in `storage` table matching by id.
func (d *Database) Update(record Record) {
	if !d.exists(record.Id) {
		return
	}

	encryptedRecord := record.Encrypt(d.encryptor)

	_, err := d.database.Exec(
		`
		UPDATE storage
		SET data = ?
		WHERE id = ?
		`,
		string(encryptedRecord.Data),
		encryptedRecord.id,
	)

	if err != nil {
		log.Fatalf("Update error: %v", err)
	}
}

// Delete record from `storage` table matching by id.
func (d *Database) Delete(record Record) {
	if !d.exists(record.Id) {
		return
	}

	_, err := d.database.Exec(
		`
		DELETE FROM storage
		WHERE id = ?;
		`,
		record.Id,
	)

	if err != nil {
		log.Fatalf("Delete error: %v", err)
	}
}
