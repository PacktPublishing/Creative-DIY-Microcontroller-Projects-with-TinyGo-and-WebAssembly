package main

import (
	"machine"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter03/keypad"
)

func main() {
	keypadDevice := keypad.Driver{}
	keypadDevice.Configure(machine.D3, machine.D4, machine.D5, machine.D6, machine.D7, machine.D8, machine.D9, machine.D10)

	for {
		key := keypadDevice.GetKey()
		if key != "" {
			println("Button: ", key)
		}
	}
}
