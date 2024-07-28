package exporters

import (
	"log"
	"strings"

	"github.com/hikkiyomi/passman/internal/exporters/mappers"
)

func GetExporter(exporter, path, browser string) Exporter {
	var result Exporter

	switch {
	case exporter == "csv" || strings.HasSuffix(path, ".csv"):
		result = NewCsvExporter(path, mappers.GetMapper(browser))
	case exporter == "tsv" || strings.HasSuffix(path, ".tsv"):
		result = NewTsvExporter(path, mappers.GetMapper(browser))
	case exporter == "json" || strings.HasSuffix(path, ".json"):
		result = NewJsonExporter(path)
	default:
		log.Fatal("Could not get exporter by given type.")
	}

	return result
}
