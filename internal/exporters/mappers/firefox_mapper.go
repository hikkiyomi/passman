package mappers

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/exporters/browser"
	"github.com/spf13/viper"
)

type firefoxMapper struct {
}

func NewFirefoxMapper() firefoxMapper {
	return firefoxMapper{}
}

func (m firefoxMapper) MapToRecord(inputData any) databases.Record {
	data, ok := inputData.([]string)
	if !ok {
		log.Fatal("Couldn't convert input data into []string.")
	}

	info := browser.FirefoxInfo{
		Url:                 data[0],
		Username:            data[1],
		Password:            data[2],
		HttpRealm:           data[3],
		FormActionOrigin:    data[4],
		Guid:                data[5],
		TimeCreated:         data[6],
		TimeLastUsed:        data[7],
		TimePasswordChanged: data[8],
	}

	return databases.Record{
		Owner:   viper.Get("user").(string),
		Service: info.Url,
		Data:    info.GetData(),
	}
}
