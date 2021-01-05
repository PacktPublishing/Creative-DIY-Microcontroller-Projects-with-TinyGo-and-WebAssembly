package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch4/buzzer"
)

func main() {
	buzzer := buzzer.NewBuzzer(machine.PD3)
	buzzer.Configure()

	for {
		buzzer.Beep(time.Millisecond*20, 3)
	}
}
