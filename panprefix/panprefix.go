package panprefix

import (
	"encoding/json"
	"log"
	"tectonic_cards/config"
)

var panPrefix string

type CardPrefix struct {
	PrefixRangeIdentifier string `json:"prefixRangeIdentifier"`
	PrefixStartRange      string `json:"prefixStartRange"`
}

type Resource struct {
	CardPrefixes []CardPrefix `json:"cardPrefixes"`
}

type ResponseBody struct {
	Resource Resource `json:"resource"`
}

func RetrievePANPrefix() string {
	body, _ := config.MakeHTTPRequest("GET", "/programadmin/v1/sponsors/configuration", nil)
	var responseBody ResponseBody
	err := json.Unmarshal([]byte(body), &responseBody)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON response: %v", err)
	}
	for _, cardPrefix := range responseBody.Resource.CardPrefixes {
		if cardPrefix.PrefixRangeIdentifier == "PAN" && len(cardPrefix.PrefixStartRange) == 12 {
			panPrefix = cardPrefix.PrefixStartRange
			break
		}
	}
	return panPrefix
}
