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

func PerformCardTransaction(panPrefix, dateTimeLocal, howPresented string, isDomestic bool) (string, error) {
	url := "/validation/v1/decisions"
	merchantInfo := MerchantInfo{
		CountryCode:          "GBR",
		CurrencyCode:         "826",
		MerchantCategoryCode: "5813",
	}
	if !isDomestic {
		merchantInfo.AddressLines = []string{"221B Baker St"}
		merchantInfo.CardAcceptorTerminalID = "1"
		merchantInfo.City = "London"
		merchantInfo.Name = "Holmes Detective Agency"
		merchantInfo.TransactionAmount = 100
	}
	payload := CardPresentTransactionPayload{
		PrimaryAccountNumber: fmt.Sprintf("%s0001", panPrefix),
		DecisionType:         "COMPLETE",
		CardholderBillAmount: 50,
		MerchantInfo:         merchantInfo,
		MessageType:          "0100",
		PointOfServiceInfo: PointOfServiceInfo{
			PersonalIdentificationNumberEntryMode: "UNKNOWN",
			PresentationData: PresentationData{
				HowPresented:  howPresented,
				IsCardPresent: howPresented == "CUSTOMER_PRESENT",
			},
			PrimaryAccountNumberEntryMode: "MAG_STRIPE_READ",
			SecurityCondition:             "NO_SECURITY_CONCERN",
			TerminalClass: TerminalClass{
				DeviceLocation: "ON_PREMISE",
				HowOperated:    "CUSTOMER_OPERATED",
				IsAttended:     true,
			},
			TerminalEntryCapability: "MAG_STRIPE_READ",
			TerminalType:            "POS_TERMINAL",
		},
		ProcessingCode:           "000000",
		RetrievalReferenceNumber: "R00000001",
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

func RetrieveControls(panPrefix string) (string, error) {
	url := "/customerrules/v1/transactiontypecontrols/cardinquiry"
	payload := CardInquiryPayload{
		PrimaryAccountNumber: fmt.Sprintf("%s0001", panPrefix),
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
