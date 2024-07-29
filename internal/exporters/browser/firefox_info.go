package browser

import (
	"encoding/json"
	"log"
)

type FirefoxInfo struct {
	Url                 string `json:"url"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	HttpRealm           string `json:"httpRealm"`
	FormActionOrigin    string `json:"formActionOrigin"`
	Guid                string `json:"guid"`
	TimeCreated         string `json:"timeCreated"`
	TimeLastUsed        string `json:"timeLastUsed"`
	TimePasswordChanged string `json:"timePasswordChanged"`
}

func (i FirefoxInfo) GetData() []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Error while marshalling firefox info: %v", err)
	}

	return data
}
