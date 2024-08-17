package mappers

import "github.com/hikkiyomi/passman/internal/databases"

type Mapper interface {
	MapToRecord(any) databases.Record
}
