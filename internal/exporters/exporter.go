package exporters

import "github.com/hikkiyomi/passman/internal/databases"

type Exporter interface {
	Import() []databases.Record
	Export([]databases.Record)
}
