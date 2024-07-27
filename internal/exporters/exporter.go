package exporters

import "github.com/hikkiyomi/passman/internal/databases"

type Exporter interface {
	Import(string) []databases.Record
	Export([]databases.Record) []string
}
