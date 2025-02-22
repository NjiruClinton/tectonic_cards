package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	//pan := panprefix.RetrievePANPrefix()
	//prefix := "0001"
	//println(pan)
	//registered, _ := registercard.RegisterCard(pan)
	//println("registered with ID: ", registered)
	//turnoff, _ := registercard.ToggleCard(registered, false)
	//println("turnoff response: ", turnoff)
	//turnon, _ := registercard.ToggleCard(registered, true)
	//println("turnon response: ", turnon)
	////Perform a Card Present transaction
	////DateTime Local should be in format: MMddhhmmss
	//data := map[string]interface{}{
	//	"countryCode":                           "GBR",
	//	"currencyCode":                          "826",
	//	"merchantCategoryCode":                  "5813",
	//	"addressLines":                          []string{"221B Baker St"},
	//	"cardAcceptorTerminalID":                "1",
	//	"city":                                  "London",
	//	"name":                                  "Holmes Detective Agency",
	//	"transactionAmount":                     100.0,
	//	"cardholderBillAmount":                  50.0,
	//	"personalIdentificationNumberEntryMode": "UNKNOWN",
	//	"primaryAccountNumberEntryMode":         "MAG_STRIPE_READ",
	//	"securityCondition":                     "NO_SECURITY_CONCERN",
	//	"deviceLocation":                        "ON_PREMISE",
	//	"howOperated":                           "CUSTOMER_OPERATED",
	//	"isAttended":                            true,
	//	"terminalEntryCapability":               "MAG_STRIPE_READ",
	//	"terminalType":                          "POS_TERMINAL",
	//	"processingCode":                        "000000",
	//	"retrievalReferenceNumber":              "R00000001",
	//}
	//present, _ := transactions.PerformCardTransaction(pan, "0101123456", "CUSTOMER_PRESENT", true, data)
	//notPresent, _ := transactions.PerformCardTransaction(pan, "0101123456", "CUSTOMER_NOT_PRESENT", false, data)
	//println("Card Present transaction response: ", present)
	//println("Card Not Present transaction response: ", notPresent)
	//deleteCard, _ := registercard.DeleteCard(registered)
	//println("deleteCard response: ", deleteCard)
	//retriveControls, _ := transactions.RetrieveControls(pan, prefix)
	//println("retriveControls response: ", retriveControls)
	//createCustomer, _ := customer.CreateCustomer(pan, "njiruclinton56@gmail.com", "Clinton", "Njiru")
	//println("createCustomer response: ", createCustomer)

}
