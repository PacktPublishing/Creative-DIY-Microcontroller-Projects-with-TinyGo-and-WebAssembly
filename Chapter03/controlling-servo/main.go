package main

import (
	"machine"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter03/servo"
)

func main() {
	servo := servo.Driver{}
	servo.Configure(machine.D11)

	servo.Right()
}
