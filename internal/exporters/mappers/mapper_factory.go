package mappers

func GetMapper(browser string) Mapper {
	var result Mapper

	switch browser {
	case "chrome":
		result = NewChromeMapper()
	case "firefox":
		result = NewFirefoxMapper()
	default:
		result = NewDefaultCsvMapper()
	}

	return result
}
