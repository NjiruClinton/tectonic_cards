package registercard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tectonic_cards/config"
)

type GlobalControl struct {
	ShouldDeclineAll     bool `json:"shouldDeclineAll"`
	ShouldAlertOnDecline bool `json:"shouldAlertOnDecline"`
	IsControlEnabled     bool `json:"isControlEnabled"`
}

type ControlPayload struct {
	GlobalControls []GlobalControl `json:"globalControls"`
}

func ToggleCard(docId string, turnOn bool) (string, error) {
	url := fmt.Sprintf("/customerrules/v1/consumertransactioncontrols/%s/rules", docId)
	payload := ControlPayload{
		GlobalControls: []GlobalControl{
			{
				ShouldDeclineAll:     !turnOn,
				ShouldAlertOnDecline: false,
				IsControlEnabled:     true,
			},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	body, err := config.MakeHTTPRequest("PUT", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return "", err
	}
	return body, nil
}
