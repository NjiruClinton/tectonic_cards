package customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tectonic_cards/config"
)

type AlertPreference struct {
	ContactType          string `json:"contactType"`
	ContactValue         string `json:"contactValue"`
	IsVerified           bool   `json:"isVerified"`
	Status               string `json:"status"`
	PreferredEmailFormat string `json:"preferredEmailFormat"`
}

type Payload struct {
	CountryCode              string            `json:"countryCode"`
	DefaultAlertsPreferences []AlertPreference `json:"defaultAlertsPreferences"`
	FirstName                string            `json:"firstName"`
	LastName                 string            `json:"lastName"`
	PreferredLanguage        string            `json:"preferredLanguage"`
	PrimaryAccountNumber     string            `json:"primaryAccountNumber"`
	UserIdentifier           string            `json:"userIdentifier"`
}

func CreateCustomer(panPrefix, email, firstName, lastName string) (string, error) {
	url := "/customerrules/v1/consumertransactioncontrols/customer"
	payload := Payload{
		CountryCode: "USA",
		DefaultAlertsPreferences: []AlertPreference{
			{
				ContactType:          "Email",
				ContactValue:         email,
				IsVerified:           true,
				Status:               "Active",
				PreferredEmailFormat: "Html",
			},
		},
		FirstName:            firstName,
		LastName:             lastName,
		PreferredLanguage:    "en-us",
		PrimaryAccountNumber: fmt.Sprintf("%s0001", panPrefix),
		UserIdentifier:       fmt.Sprintf("%s%s-guid", firstName, lastName),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	body, err := config.MakeHTTPRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return "", err
	}
	return body, nil
}
