package main

import (
	"machine"
	"time"
)

var stopTraffic bool

// It is important to set a scheduler for this to work
// tinygo flash -scheduler tasks -target=arduino ch2/traffic-lights-pedestrian/main.go
func main() {
	stopTraffic = false

	outputConfig := machine.PinConfig{Mode: machine.PinOutput}

	greenLED := machine.D11
	greenLED.Configure(outputConfig)

	yellowLED := machine.D12
	yellowLED.Configure(outputConfig)

	redLED := machine.D13
	redLED.Configure(outputConfig)

	pedestrianGreen := machine.D5
	pedestrianGreen.Configure(outputConfig)
	pedestrianRed := machine.D4
	pedestrianRed.Configure(outputConfig)

	inputConfig := machine.PinConfig{Mode: machine.PinInput}
	buttonInput := machine.D2
	buttonInput.Configure(inputConfig)

	pedestrianRed.High()
	pedestrianGreen.Low()

	go trafficLights(redLED, greenLED, yellowLED, pedestrianRed, pedestrianGreen)

	for {
		if buttonInput.Get() {
			stopTraffic = true
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func trafficLights(redLED, greenLED, yellowLED, pedestrianRED, pedestrianGreen machine.Pin) {
	for {
		if stopTraffic {
			redLED.High()
			yellowLED.Low()
			greenLED.Low()

			pedestrianGreen.High()
			pedestrianRED.Low()
			time.Sleep(3 * time.Second)

			pedestrianGreen.Low()
			pedestrianRED.High()

			stopTraffic = false
		} else {
			pedestrianGreen.Low()
			pedestrianRED.High()
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
