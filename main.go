package main

import (
	"tectonic_cards/panprefix"
)

func main() {
	pan := panprefix.RetrievePANPrefix()
	println(pan)
	//registercard.RegisterCard(pan)

}
