package transactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tectonic_cards/config"
)

type MerchantInfo struct {
	AddressLines           []string `json:"addressLines,omitempty"`
	CardAcceptorTerminalID string   `json:"cardAcceptorTerminalID,omitempty"`
	City                   string   `json:"city,omitempty"`
	CountryCode            string   `json:"countryCode"`
	CurrencyCode           string   `json:"currencyCode"`
	MerchantCategoryCode   string   `json:"merchantCategoryCode"`
	Name                   string   `json:"name,omitempty"`
	TransactionAmount      int      `json:"transactionAmount,omitempty"`
}

type PresentationData struct {
	HowPresented  string `json:"howPresented"`
	IsCardPresent bool   `json:"isCardPresent"`
}

type TerminalClass struct {
	DeviceLocation string `json:"deviceLocation"`
	HowOperated    string `json:"howOperated"`
	IsAttended     bool   `json:"isAttended"`
}

type PointOfServiceInfo struct {
	PersonalIdentificationNumberEntryMode string           `json:"personalIdentificationNumberEntryMode"`
	PresentationData                      PresentationData `json:"presentationData"`
	PrimaryAccountNumberEntryMode         string           `json:"primaryAccountNumberEntryMode"`
	SecurityCondition                     string           `json:"securityCondition"`
	TerminalClass                         TerminalClass    `json:"terminalClass"`
	TerminalEntryCapability               string           `json:"terminalEntryCapability"`
	TerminalType                          string           `json:"terminalType"`
}

type CardPresentTransactionPayload struct {
	PrimaryAccountNumber     string             `json:"primaryAccountNumber"`
	DecisionType             string             `json:"decisionType"`
	CardholderBillAmount     int                `json:"cardholderBillAmount"`
	MerchantInfo             MerchantInfo       `json:"merchantInfo"`
	MessageType              string             `json:"messageType"`
	PointOfServiceInfo       PointOfServiceInfo `json:"pointOfServiceInfo"`
	ProcessingCode           string             `json:"processingCode"`
	RetrievalReferenceNumber string             `json:"retrievalReferenceNumber"`
	DateTimeLocal            string             `json:"dateTimeLocal"`
}

type CardInquiryPayload struct {
	PrimaryAccountNumber string `json:"primaryAccountNumber"`
}

func PerformCardTransaction(panPrefix, dateTimeLocal, howPresented string, isDomestic bool, data map[string]interface{}) (string, error) {
	url := "/validation/v1/decisions"
	merchantInfo := MerchantInfo{
		CountryCode:          data["countryCode"].(string),
		CurrencyCode:         data["currencyCode"].(string),
		MerchantCategoryCode: data["merchantCategoryCode"].(string),
	}
	if !isDomestic {
		merchantInfo.AddressLines = data["addressLines"].([]string)
		merchantInfo.CardAcceptorTerminalID = data["cardAcceptorTerminalID"].(string)
		merchantInfo.City = data["city"].(string)
		merchantInfo.Name = data["name"].(string)
		merchantInfo.TransactionAmount = int(data["transactionAmount"].(float64))
	}
	payload := CardPresentTransactionPayload{
		PrimaryAccountNumber: fmt.Sprintf("%s0001", panPrefix),
		DecisionType:         "COMPLETE",
		CardholderBillAmount: int(data["cardholderBillAmount"].(float64)),
		MerchantInfo:         merchantInfo,
		MessageType:          "0100",
		PointOfServiceInfo: PointOfServiceInfo{
			PersonalIdentificationNumberEntryMode: data["personalIdentificationNumberEntryMode"].(string),
			PresentationData: PresentationData{
				HowPresented:  howPresented,
				IsCardPresent: howPresented == "CUSTOMER_PRESENT",
			},
			PrimaryAccountNumberEntryMode: data["primaryAccountNumberEntryMode"].(string),
			SecurityCondition:             data["securityCondition"].(string),
			TerminalClass: TerminalClass{
				DeviceLocation: data["deviceLocation"].(string),
				HowOperated:    data["howOperated"].(string),
				IsAttended:     data["isAttended"].(bool),
			},
			TerminalEntryCapability: data["terminalEntryCapability"].(string),
			TerminalType:            data["terminalType"].(string),
		},
		ProcessingCode:           data["processingCode"].(string),
		RetrievalReferenceNumber: data["retrievalReferenceNumber"].(string),
		DateTimeLocal:            dateTimeLocal,
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

func RetrieveControls(pan, prefix string) (string, error) {
	url := "/customerrules/v1/transactiontypecontrols/cardinquiry"
	payload := CardInquiryPayload{
		PrimaryAccountNumber: pan + prefix,
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
