package mappers

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/spf13/viper"
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
		Owner:   viper.Get("user").(string),
		Service: data[1],
		Data:    []byte(data[2]),
	}
}
