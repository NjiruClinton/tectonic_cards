package panprefix

import (
	"encoding/json"
	"log"
	"net/http"
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
	USER_ID, PASSWORD := config.LoadEnvVariables()
	AUTH := config.GetAuthHeader(USER_ID, PASSWORD)
	BASIC_URL := "https://sandbox.api.visa.com/vctc"

	tlsConfig, err := config.SetupTLSConfig("./cert.pem", "./key.pem", "./cacert.pem")
	if err != nil {
		log.Fatalf("Error setting up TLS configuration: %v", err)
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	url := BASIC_URL + "/programadmin/v1/sponsors/configuration"
	body, _ := config.MakeHTTPRequest(client, url, AUTH)
	var responseBody ResponseBody
	err = json.Unmarshal([]byte(body), &responseBody)
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
