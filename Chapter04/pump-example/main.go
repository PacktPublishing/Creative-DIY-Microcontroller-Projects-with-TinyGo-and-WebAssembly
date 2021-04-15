package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter04/pump"
)

func main() {
	pump := pump.NewPump(machine.D5)
	pump.Configure()

	for {
		pump.Pump(350*time.Millisecond, 3)
		time.Sleep(30 * time.Second)
	}
}
