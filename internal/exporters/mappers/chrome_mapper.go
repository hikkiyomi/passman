package mappers

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/exporters/browser"
	"github.com/spf13/viper"
)

type chromeMapper struct {
}

func NewChromeMapper() chromeMapper {
	return chromeMapper{}
}

func (m chromeMapper) MapToRecord(inputData any) databases.Record {
	data, ok := inputData.([]string)
	if !ok {
		log.Fatal("Couldn't convert input data into []string.")
	}

	info := browser.ChromeInfo{
		Name:     data[0],
		Url:      data[1],
		Username: data[2],
		Password: data[3],
		Note:     data[4],
	}

	return databases.Record{
		Owner:   viper.Get("user").(string),
		Service: info.Name,
		Data:    info.GetData(),
	}
}
