package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter03/keypad"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter03/servo"
)

func main() {
	keypadDriver := keypad.Driver{}
	keypadDriver.Configure(machine.D2, machine.D3, machine.D4, machine.D5, machine.D6, machine.D7, machine.D8, machine.D9)

	servoDriver := servo.Driver{}
	servoDriver.Configure(machine.D11)

	outPutConfig := machine.PinConfig{Mode: machine.PinOutput}

	led1 := machine.D12
	led1.Configure(outPutConfig)

	led2 := machine.D13
	led2.Configure(outPutConfig)

	const passcode = "133742"
	enteredPasscode := ""

	for {
		key := keypadDriver.GetKey()
		if key != "" {
			println("Button: ", key)

			enteredPasscode += key

			led2.High()
			time.Sleep(time.Second / 5)
			led2.Low()
		}

		if len(enteredPasscode) == len(passcode) {
			if enteredPasscode == passcode {
				println("Success")
				enteredPasscode = ""
				servoDriver.Right()

				led1.High()
				time.Sleep(time.Second * 3)
				led1.Low()

			} else {
				println("Fail")
				println("Entered Password: ", enteredPasscode)
				enteredPasscode = ""

				led2.High()
				time.Sleep(time.Second * 3)
				led2.Low()
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}
