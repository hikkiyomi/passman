package exporters

import "log"

func GetExporterByType(exporter, path string) Exporter {
	var result Exporter

	switch exporter {
	case "csv":
		result = NewCsvExporter(path)
	case "tsv":
		result = NewTsvExporter(path)
	case "json":
		result = NewJsonExporter(path)
	default:
		log.Fatal("Could not get exporter by given type.")
	}

	return result
}
