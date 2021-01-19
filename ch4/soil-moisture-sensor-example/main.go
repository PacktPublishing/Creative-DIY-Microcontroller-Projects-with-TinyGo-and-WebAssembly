package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch4/soil-moisture-sensor"
)

func main() {
	machine.InitADC()
	soilSensor := soil.NewSoilSensor(17600, 37888, machine.ADC5, machine.D2)

	for {
		soilSensor.On()
		time.Sleep(100 * time.Millisecond)
		println("current humidity: ", soilSensor.Get(), "%")
		soilSensor.Off()
		time.Sleep(1 * time.Second)
	}
}
