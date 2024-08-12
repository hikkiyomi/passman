package exporters

import (
	"errors"
	"strings"

	"github.com/hikkiyomi/passman/internal/exporters/mappers"
)

func addExtensionToPath(path, ext string) string {
	if strings.HasSuffix(path, ext) {
		return path
	}

	return path + "." + ext
}

func GetExporter(exporter, path, browser string) (Exporter, error) {
	var result Exporter

	switch {
	case exporter == "csv" || strings.HasSuffix(path, ".csv"):
		result = NewCsvExporter(addExtensionToPath(path, "csv"), mappers.GetMapper(browser))
	case exporter == "tsv" || strings.HasSuffix(path, ".tsv"):
		result = NewTsvExporter(addExtensionToPath(path, "tsv"), mappers.GetMapper(browser))
	case exporter == "json" || strings.HasSuffix(path, ".json"):
		result = NewJsonExporter(addExtensionToPath(path, "json"))
	default:
		return nil, errors.New("couldn't get an exporter of given type")
	}

	return result, nil
}
