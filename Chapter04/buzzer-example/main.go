package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter04/buzzer"
)

func main() {
	buzzer := buzzer.NewBuzzer(machine.D4)
	buzzer.Configure()

	for {
		buzzer.Beep(time.Millisecond*100, 3)
		time.Sleep(3 * time.Second)
	}
}
