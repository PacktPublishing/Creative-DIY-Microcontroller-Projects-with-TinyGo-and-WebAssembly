package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter04/buzzer"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter04/pump"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter04/soil-moisture-sensor"
	waterlevel "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter04/water-level-sensor"
)

func main() {
	machine.InitADC()

	soilSensor := soil.NewSoilSensor(18000, 34800, machine.ADC5, machine.D2)
	soilSensor.Configure()

	waterLevelSensor := waterlevel.NewWaterLevel(7000, machine.ADC4, machine.D3)
	waterLevelSensor.Configure()

	pump := pump.NewPump(machine.D5)
	pump.Configure()

	buzzer := buzzer.NewBuzzer(machine.D4)
	buzzer.Configure()

	for {
		waterLevelSensor.On()
		time.Sleep(100 * time.Millisecond)
		if waterLevelSensor.IsEmpty() {
			waterLevelSensor.Off()
			buzzer.Beep(150*time.Millisecond, 3)
			time.Sleep(time.Hour)
			continue
		}
		waterLevelSensor.Off()

		soilSensor.On()
		time.Sleep(100 * time.Millisecond)
		switch soilSensor.Get() {
		case soil.CompletelyDry:
			fallthrough
		case soil.VeryDry:
			soilSensor.Off()
			pump.Pump(350*time.Millisecond, 3)
		default:
			soilSensor.Off()
			time.Sleep(time.Hour)
		}
	}
}
