package main

import (
	"machine"
	"time"
)

func main() {
	machine.InitADC()
	soilSensor := machine.ADC{Pin: machine.ADC5}
	soilSensor.Configure(machine.ADCConfig{})

	machine.D2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.D2.High()

	for {
		value := soilSensor.Get()
		println(value)
		time.Sleep(500 * time.Millisecond)
	}
}
