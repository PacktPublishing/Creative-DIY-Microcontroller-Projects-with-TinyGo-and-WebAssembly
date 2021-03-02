package main

import (
	"machine"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter03/keypad"
)

func main() {
	keypadDevice := keypad.Driver{}
	keypadDevice.Configure(machine.D2, machine.D3, machine.D4, machine.D5, machine.D6, machine.D7, machine.D8, machine.D9)

	for {
		key := keypadDevice.GetKey()
		if key != "" {
			println("Button: ", key)
		}
	}
}
