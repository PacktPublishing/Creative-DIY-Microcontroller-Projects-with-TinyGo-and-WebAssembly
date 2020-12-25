package main

import "machine"

func main() {
	machine.InitADC()
	soilSensor := machine.ADC{Pin: machine.ADC3}
	soilSensor.Configure()

	// AVR has 10 bit precision
	value := soilSensor.Get()
	println(value)
}
