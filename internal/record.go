package internal

import "database/sql"

type Record struct {
	Owner   string
	Service string
	Data    []byte
}

func (r *Record) Scan(rows *sql.Rows) error {
	var data string

	err := rows.Scan(&r.Owner, &r.Service, &data)
	r.Data = []byte(data)

	return err
}
