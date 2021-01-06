package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch4/buzzer"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch4/pump"
)

func main() {
	machine.InitADC()
	soilSensor := machine.ADC{Pin: machine.ADC3}
	soilSensor.Configure()

	waterLevelSensor := machine.ADC{Pin: machine.ADC5}
	waterLevelSensor.Configure()

	buzzer := buzzer.NewBuzzer(machine.D3)
	buzzer.Configure()

	pump := pump.NewPump(machine.D8)
	pump.Configure()

	go checkWaterLevel(waterLevelSensor, buzzer)

	for {
		soilLevel := soilSensor.Get()
		if soilLevel < 16.000 {
			pump.Pump(300*time.Millisecond, 3)
		}

		time.Sleep(time.Minute)
	}
}

func checkWaterLevel(waterLevelSensor machine.ADC, buzzer buzzer.Buzzer) {
	for {
		waterLevel := waterLevelSensor.Get()
		if waterLevel < 25.000 {
			buzzer.Beep(time.Millisecond*100, 3)
		}

		time.Sleep(time.Hour)
	}
}
