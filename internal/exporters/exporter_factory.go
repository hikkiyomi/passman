package exporters

import (
	"log"

	"github.com/hikkiyomi/passman/internal/exporters/mappers"
)

func GetExporter(exporter, path, browser string) Exporter {
	var result Exporter

	switch exporter {
	case "csv":
		result = NewCsvExporter(path, mappers.GetMapper(browser))
	case "tsv":
		result = NewTsvExporter(path, mappers.GetMapper(browser))
	case "json":
		result = NewJsonExporter(path)
	default:
		log.Fatal("Could not get exporter by given type.")
	}

	return result
}
