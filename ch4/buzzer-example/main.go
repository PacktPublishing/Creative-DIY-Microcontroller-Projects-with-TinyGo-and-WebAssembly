package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch4/buzzer"
)

func main() {
	buzzer := buzzer.NewBuzzer(machine.D4)
	buzzer.Configure()

	for {
		buzzer.Beep(time.Millisecond*100, 3)
		time.Sleep(3 * time.Second)
	}
}
