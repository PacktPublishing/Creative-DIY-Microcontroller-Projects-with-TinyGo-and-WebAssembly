package main

import (
	"machine"
	"time"

	waterlevel "github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter04/water-level-sensor"
)

func main() {
	machine.InitADC()

	waterLevelSensor := waterlevel.NewWaterLevel(7000, machine.ADC4, machine.D3)
	waterLevelSensor.Configure()

	for {
		waterLevelSensor.On()
		time.Sleep(100 * time.Millisecond)
		println("tank is empty", waterLevelSensor.IsEmpty())
		waterLevelSensor.Off()
		time.Sleep(time.Second)
	}
}
