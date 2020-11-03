package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}

	redLED := machine.D13
	redLED.Configure(outputConfig)

	for {
		redLED.Low()
		time.Sleep(500 * time.Millisecond)
		redLED.High()
		time.Sleep(500 * time.Millisecond)
	}
}
