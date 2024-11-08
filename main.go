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
	turnoff, _ := registercard.ToggleCard(registered, false)
	println("turnoff response: ", turnoff)
	turnon, _ := registercard.ToggleCard(registered, true)
	println("turnon response: ", turnon)

}
