package exporters

import "github.com/hikkiyomi/passman/internal/databases"

// TSV is just CSV with tab separator.
type tsvExporter struct {
	csvExporter csvExporter
}

func NewTsvExporter(path string) tsvExporter {
	csvExporter := NewCsvExporter(path)
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
