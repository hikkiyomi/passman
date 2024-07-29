package exporters

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/exporters/mappers"
)

type csvExporter struct {
	comma  rune
	path   string
	mapper mappers.Mapper
}

func NewCsvExporter(path string, mapper mappers.Mapper) csvExporter {
	return csvExporter{
		comma:  ',',
		path:   path,
		mapper: mapper,
	}
}

func (e csvExporter) Import() []databases.Record {
	f, err := os.Open(e.path)
	if err != nil {
		log.Fatalf("Error while opening file for importing csv: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = e.comma

	result := make([]databases.Record, 0)
	haveReadColumnNames := false

	for {
		data, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading csv: %v", err)
		}

		if !haveReadColumnNames {
			haveReadColumnNames = true
			continue
		}

		result = append(result, e.mapper.MapToRecord(data))
	}

	return result
}

func (e csvExporter) Export(records []databases.Record) {
	data := [][]string{
		{"owner", "service", "data"},
	}

	for _, record := range records {
		data = append(data, []string{record.Owner, record.Service, string(record.Data)})
	}

	f, err := os.Create(e.path)
	if err != nil {
		log.Fatalf("Error while creating file for exporting csv: %v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = e.comma
	defer w.Flush()

	for _, line := range data {
		if err := w.Write(line); err != nil {
			log.Fatalf("Error while writing to csv file: %v", err)
		}
	}
}
