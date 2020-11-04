package main

import (
	"machine"
)

// This example is used to show floating of the buttons value, when no Pull-Up resistor is being used
// The button state is printed to serial, in order to watch the effect
func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	inputConfig := machine.PinConfig{Mode: machine.PinInput}

	// To get around this PinInputPullUp can be used. If the pin is configured with PinInputPullup an internal resistor is being used.
	// inputConfig := machine.PinConfig{Mode: machine.PinInputPullup}

	led := machine.D13
	led.Configure(outputConfig)

	buttonInput := machine.D2
	buttonInput.Configure(inputConfig)

	for {
		buttonState := buttonInput.Get()
		println(buttonState)

		if buttonState {
			led.High()
			continue
		}

		led.Low()
	}
}
