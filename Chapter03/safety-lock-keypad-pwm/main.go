package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter03/keypad"
	servopwm "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter03/servo-pwm"
)

func main() {
	keypadDriver := keypad.Driver{}
	// As we use the pin D9 for PWM, we also switch a pin here. D9 -> D10
	keypadDriver.Configure(machine.D2, machine.D3, machine.D4, machine.D5, machine.D6, machine.D7, machine.D8, machine.D10)

	// we are using machine.D9 here, as we want to make use of the 16 bit timer
	servoDriver := servopwm.NewDevice(machine.Timer1, machine.D9)
	servoDriver.Configure()

	outPutConfig := machine.PinConfig{Mode: machine.PinOutput}

	led1 := machine.D12
	led1.Configure(outPutConfig)

	led2 := machine.D13
	led2.Configure(outPutConfig)

	const passcode = "133742"
	enteredPasscode := ""

	servoDriver.Left()

	for {
		key := keypadDriver.GetKey()
		if key == "" {
			continue
		}

		println("Button: ", key)

		enteredPasscode += key

		led2.High()
		time.Sleep(200 * time.Millisecond)
		led2.Low()

		if len(enteredPasscode) != len(passcode) {
			continue
		}

		if enteredPasscode == passcode {
			println("Success")
			servoDriver.Right()

			led1.High()
			time.Sleep(time.Second * 3)
			led1.Low()

			time.Sleep(time.Second)
			servoDriver.Left()

		} else {
			println("Fail")
			println("Entered Password: ", enteredPasscode)

			led2.High()
			time.Sleep(time.Second * 3)
			led2.Low()
		}

		enteredPasscode = ""
		time.Sleep(50 * time.Millisecond)
	}
}
