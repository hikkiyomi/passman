package exporters

import (
	"log"
	"strings"

	"github.com/hikkiyomi/passman/internal/exporters/mappers"
)

func addExtensionToPath(path, ext string) string {
	if strings.HasSuffix(path, ext) {
		return path
	}

	return path + "." + ext
}

func GetExporter(exporter, path, browser string) Exporter {
	var result Exporter

	switch {
	case exporter == "csv" || strings.HasSuffix(path, ".csv"):
		result = NewCsvExporter(addExtensionToPath(path, "csv"), mappers.GetMapper(browser))
	case exporter == "tsv" || strings.HasSuffix(path, ".tsv"):
		result = NewTsvExporter(addExtensionToPath(path, "tsv"), mappers.GetMapper(browser))
	case exporter == "json" || strings.HasSuffix(path, ".json"):
		result = NewJsonExporter(addExtensionToPath(path, "json"))
	default:
		log.Fatal("Could not get exporter by given type.")
	}

	return result
}
