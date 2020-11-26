package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch3/keypad"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch3/servo"
)

func main() {
	keypadDriver := keypad.Driver{}
	keypadDriver.Configure(machine.D2, machine.D3, machine.D4, machine.D5, machine.D6, machine.D7, machine.D8, machine.D9)

	servoDriver := servo.Driver{}
	servoDriver.Configure(machine.D11)

	const password = "133742"
	enteredPassword := ""

	for {
		key := keypadDriver.GetKey()
		if key != "" {
			println("Button: ", key)

			enteredPassword += key
		}

		if len(enteredPassword) == len(password) {
			if password == enteredPassword {
				println("Success")
				enteredPassword = ""
				servoDriver.Right()
			} else {
				println("Fail")
				println("Entered Password: ", enteredPassword)
				enteredPassword = ""
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}
