package main

import (
	"machine"
	"time"

	servopwm "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter03/servo-pwm"
)

func main() {
	servo := servopwm.NewDevice(machine.Timer1, machine.D9)
	err := servo.Configure()
	if err != nil {
		for {
			println("could not configure servo:", err.Error())
			time.Sleep(time.Second)
		}
	}

	for {
		servo.Left()
		time.Sleep(time.Second)

		servo.Center()
		time.Sleep(time.Second)

		servo.Right()
		time.Sleep(time.Second)
	}
}
