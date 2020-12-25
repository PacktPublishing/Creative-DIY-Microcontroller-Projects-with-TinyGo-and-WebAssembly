package main

import (
	"machine"
)

func main() {
	machine.InitADC()
	waterLevelSensor := machine.ADC{Pin: machine.ADC5}
	waterLevelSensor.Configure()

	// AVR has 10 bit precision
	value := waterLevelSensor.Get()
	println(value)
}
