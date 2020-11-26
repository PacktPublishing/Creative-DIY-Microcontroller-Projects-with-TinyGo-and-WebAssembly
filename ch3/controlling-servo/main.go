package main

import (
	"machine"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch3/servo"
)

func main() {
	servo := servo.Servo{}
	servo.Configure(machine.D11)

	servo.Right()
}
