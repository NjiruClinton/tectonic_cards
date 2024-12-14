package main

import (
	"tectonic_cards/panprefix"
	"tectonic_cards/registercard"
)

func main() {
	pan := panprefix.RetrievePANPrefix()
	println(pan)
	registered, _ := registercard.RegisterCard(pan)
	println("registered with ID: ", registered)
	//turnoff, _ := registercard.ToggleCard(registered, false)
	//println("turnoff response: ", turnoff)
	//turnon, _ := registercard.ToggleCard(registered, true)
	//println("turnon response: ", turnon)
	// Perform a Card Present transaction
	// DateTime Local should be in format: MMddhhmmss
	//present, _ := transactions.PerformCardTransaction(pan, "0101123456", "CUSTOMER_PRESENT", true)
	//notPresent, _ := transactions.PerformCardTransaction(pan, "0101123456", "CUSTOMER_NOT_PRESENT", false)
	//println("Card Present transaction response: ", present)
	//println("Card Not Present transaction response: ", notPresent)
	deleteCard, _ := registercard.DeleteCard(registered)
	println("deleteCard response: ", deleteCard)
	//retriveControls, _ := transactions.RetrieveControls(pan)
	//println("retriveControls response: ", retriveControls)
	//createCustomer, _ := customer.CreateCustomer(pan, "njiruclinton56@gmail.com", "Clinton", "Njiru")
	//println("createCustomer response: ", createCustomer)

}
