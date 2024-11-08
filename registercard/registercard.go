package registercard

import (
	"encoding/json"
	"fmt"
	"strings"
	"tectonic_cards/config"
)

type Response struct {
	Resource struct {
		DocumentID string `json:"documentID"`
	} `json:"resource"`
}

func RegisterCard(panPrefix string) (string, error) {
	payload := strings.NewReader(fmt.Sprintf(`{
	"primaryAccountNumber": "%s0001"
		}`, panPrefix))
	body, err := config.MakeHTTPRequest("POST", "/customerrules/v1/consumertransactioncontrols", payload)
	if err != nil {
		return "", err
	}
	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", err
	}
	return response.Resource.DocumentID, nil
}
