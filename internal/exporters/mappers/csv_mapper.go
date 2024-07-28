package mappers

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
)

type defaultCsvMapper struct {
}

func NewDefaultCsvMapper() defaultCsvMapper {
	return defaultCsvMapper{}
}

func (m defaultCsvMapper) MapToRecord(inputData any) databases.Record {
	data, ok := inputData.([]string)
	if !ok {
		log.Fatal("Couldn't convert input data into []string.")
	}

	return databases.Record{
		Owner:   data[0],
		Service: data[1],
		Data:    []byte(data[2]),
	}
}
