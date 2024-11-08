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

}
