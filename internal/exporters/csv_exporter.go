package exporters

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/hikkiyomi/passman/internal/databases"
)

type csvExporter struct {
	comma rune
	path  string
}

func NewCsvExporter(path string) csvExporter {
	return csvExporter{
		comma: ',',
		path:  path,
	}
}

func (e csvExporter) Import() []databases.Record {
	f, err := os.Open(e.path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = e.comma
	records := make([]databases.Record, 0)
	haveReadColumns := false

	for {
		data, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if !haveReadColumns {
			haveReadColumns = true
			continue
		}

		records = append(records, databases.Record{
			Owner:   data[0],
			Service: data[1],
			Data:    []byte(data[2]),
		})
	}

	return records
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
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = e.comma
	defer w.Flush()

	for _, line := range data {
		if err := w.Write(line); err != nil {
			log.Fatal(err)
		}
	}
}
