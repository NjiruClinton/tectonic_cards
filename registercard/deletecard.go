package registercard

import (
	"fmt"
	"tectonic_cards/config"
)

func DeleteCard(docId string) (string, error) {
	url := fmt.Sprintf("/customerrules/v1/consumertransactioncontrols/%s", docId)
	body, err := config.MakeHTTPRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}
	return body, nil
}
