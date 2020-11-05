package main

import (
	"machine"
	"time"
)

var blink bool

// It is important to set a scheduler for this to work
// tinygo flash -scheduler tasks -target=arduino ch2/traffic-lights-pedestrian/main.go
func main() {
	blink = false

	outputConfig := machine.PinConfig{Mode: machine.PinOutput}

	greenLED := machine.D13
	greenLED.Configure(outputConfig)

	yellowLED := machine.D12
	yellowLED.Configure(outputConfig)

	redLED := machine.D11
	redLED.Configure(outputConfig)

	pedestrianGreen := machine.D4
	pedestrianGreen.Configure(outputConfig)
	pedestrianRed := machine.D5
	pedestrianRed.Configure(outputConfig)

	inputConfig := machine.PinConfig{Mode: machine.PinInput}
	buttonInput := machine.D2
	buttonInput.Configure(inputConfig)

	pedestrianRed.High()
	pedestrianGreen.Low()

	go trafficLights(redLED, greenLED, yellowLED, pedestrianRed, pedestrianGreen)

	for {
		if buttonInput.Get() {
			blink = !blink
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func trafficLights(redLED, greenLED, yellowLED, pedestrianRED, pedestrianGreen machine.Pin) {
	for {
		if blink {
			redLED.Low()
			greenLED.Low()
			pedestrianGreen.Low()
			pedestrianRED.Low()

			yellowLED.High()
			time.Sleep(400 * time.Millisecond)
			yellowLED.Low()

			continue
		}

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
