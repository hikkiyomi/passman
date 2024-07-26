package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	database *sql.DB
	path     string
}

// Initializes a database and creates table `storage` if it does not exist.
func newDatabase(database *sql.DB, path string) *Database {
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
		log.Fatal(err)
	}

	return &Database{
		database: database,
		path:     path,
	}
}

// Creates a connection to database.
// `path` variable points to a directory where database "passman.db" reside, e.g. $HOME.
// If you want to create a new database, put your desirable path to it as an argument.
func Open(path string) *Database {
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	absPath := path + "/passman.db"
	db, err := sql.Open("sqlite3", absPath)

	if err != nil {
		log.Fatal(err)
	}

	return newDatabase(db, absPath)
}

// Inserts a record into `storage` table.
func (d *Database) Insert(record Record) {
	_, err := d.database.Exec(
		`
		INSERT INTO storage (owner, service, data)
		VALUES (?, ?, ?);
		`,
		record.Owner,
		record.Service,
		string(record.Data),
	)

	if err != nil {
		log.Fatal(err)
	}
}

// Retrieves tuples from `storage` table by owner.
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
		log.Fatal(err)
	}

	result := make([]Record, 0)

	for rows.Next() {
		var record Record

		if err := record.Scan(rows); err != nil {
			log.Fatal(err)
		}

		result = append(result, record)
	}

	return result
}

// Retrieves the only tuple from `storage` table matching by owner and service,
// or nil if it does not exist.
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

// Updates the data in `storage` table matching by owner and service.
func (d *Database) Update(record Record) {
	_, err := d.database.Exec(
		`
		UPDATE storage
		SET data = ?
		WHERE owner = ?
			AND service = ?;
		`,
		string(record.Data),
		record.Owner,
		record.Service,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// Deletes all passwords from `storage` table matching by owner and service.
func (d *Database) Delete(record Record) {
	_, err := d.database.Exec(
		`
		DELETE FROM storage
		WHERE owner = ?
			AND service = ?;
		`,
		record.Owner,
		record.Service,
	)

	if err != nil {
		log.Fatal(err)
	}
}
