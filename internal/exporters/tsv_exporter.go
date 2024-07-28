package exporters

import (
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/exporters/mappers"
)

// TSV is just CSV with a tab separator.
type tsvExporter struct {
	csvExporter csvExporter
}

func NewTsvExporter(path string, mapper mappers.Mapper) tsvExporter {
	csvExporter := NewCsvExporter(path, mapper)
	csvExporter.comma = '\t'

	return tsvExporter{
		csvExporter: csvExporter,
	}
}

func (e tsvExporter) Import() []databases.Record {
	return e.csvExporter.Import()
}

func (e tsvExporter) Export(records []databases.Record) {
	e.csvExporter.Export(records)
}
