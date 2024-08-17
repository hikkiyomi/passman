package exporters

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hikkiyomi/passman/internal/databases"
)

type jsonExporter struct {
	path string
}

func NewJsonExporter(path string) jsonExporter {
	return jsonExporter{
		path: path,
	}
}

func (e jsonExporter) Import() []databases.Record {
	bytes, err := os.ReadFile(e.path)
	if err != nil {
		log.Fatalf("Error while reading json file for import: %v", err)
	}

	result := make([]databases.Record, 0)

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Fatalf("Error while unmarshalling json file for import: %v", err)
	}

	return result
}

func (e jsonExporter) Export(records []databases.Record) {
	result, err := json.MarshalIndent(records, "", "    ")
	if err != nil {
		log.Fatalf("Error while marshalling exporting json: %v", err)
	}

	err = os.WriteFile(e.path, result, 0666)
	if err != nil {
		log.Fatalf("Error while writing exporting json: %v", err)
	}
}
