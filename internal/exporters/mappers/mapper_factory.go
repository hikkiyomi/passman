package mappers

import "strings"

func GetMapper(browser string) Mapper {
	var result Mapper

	switch strings.ToLower(browser) {
	case "chrome":
		result = NewChromeMapper()
	case "firefox":
		result = NewFirefoxMapper()
	default:
		result = NewDefaultCsvMapper()
	}

	return result
}
