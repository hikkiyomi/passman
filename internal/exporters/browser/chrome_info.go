package browser

import (
	"encoding/json"
	"log"
)

type ChromeInfo struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Note     string `json:"note"`
}

func (i ChromeInfo) GetData() []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
