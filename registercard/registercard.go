package registercard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tectonic_cards/config"
)

type RegisterCardRequest struct {
	PrimaryAccountNumber string `json:"primaryAccountNumber"`
}

func RegisterCard(panPrefix string) {
	AUTH := config.GetAuthHeader()
	BASIC_URL := "https://sandbox.api.visa.com/vctc"

	tlsConfig, _ := config.SetupTLSConfig()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	url := BASIC_URL + "/customerrules/v1/consumertransactioncontrols"
	fmt.Print("Registering card with PAN: ", panPrefix+"0001\n")
	requestBody := RegisterCardRequest{
		PrimaryAccountNumber: panPrefix + "0001",
	}
	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", AUTH)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Error: received non-200 response code: %v", res.StatusCode)
	}

	log.Println("Card registered successfully")
}
