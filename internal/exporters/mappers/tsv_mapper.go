package mappers

import "github.com/hikkiyomi/passman/internal/databases"

type defaultTsvMapper struct {
	csvMapper defaultCsvMapper
}

func NewDefaultTsvMapper() defaultTsvMapper {
	return defaultTsvMapper{
		csvMapper: NewDefaultCsvMapper(),
	}
}

func (m defaultTsvMapper) MapToRecord(inputData any) databases.Record {
	return m.csvMapper.MapToRecord(inputData)
}
