package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch4/pump"
)

func main() {
	pump := pump.NewPump(machine.D5)
	pump.Configure()

	for {
		pump.Pump(350*time.Millisecond, 3)
		time.Sleep(30 * time.Second)
	}
}
