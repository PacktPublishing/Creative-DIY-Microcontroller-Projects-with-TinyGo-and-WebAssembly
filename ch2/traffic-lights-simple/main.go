package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}

	greenLED := machine.D13
	greenLED.Configure(outputConfig)

	yellowLED := machine.D12
	yellowLED.Configure(outputConfig)

	redLED := machine.D11
	redLED.Configure(outputConfig)

	for {
		redLED.High()
		time.Sleep(time.Second)
		yellowLED.High()
		time.Sleep(time.Second)
		redLED.Low()
		yellowLED.Low()
		greenLED.High()
		time.Sleep(time.Second)

		greenLED.Low()
		yellowLED.High()
		time.Sleep(time.Second)
		yellowLED.Low()
	}
}
