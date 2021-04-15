package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter04/soil-moisture-sensor"
)

func main() {
	machine.InitADC()
	soilSensor := soil.NewSoilSensor(18000, 34000, machine.ADC5, machine.D2)
	soilSensor.Configure()

	for {
		soilSensor.On()
		time.Sleep(75 * time.Millisecond)

		switch soilSensor.Get() {
		case soil.CompletelyDry:
			println("completely dry")
		case soil.VeryDry:
			println("very dry")
		case soil.Dry:
			println("dry")
		case soil.Wet:
			println("wet")
		case soil.VeryWet:
			println("very wet")
		case soil.Water:
			println("pure water")
		}
		soilSensor.Off()

		time.Sleep(1 * time.Second)
	}
}
