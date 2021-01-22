package waterlevel

import "machine"

type WaterLevel interface {
	IsEmpty() bool
	Configure()
	On()
	Off()
}

type waterLevel struct {
	dryThreshold uint16
	pin          machine.Pin
	adc          machine.ADC
	voltage      machine.Pin
}

func NewWaterLevel(dryThreshold uint16, dataPin, voltagePin machine.Pin) WaterLevel {
	return &waterLevel{
		dryThreshold: dryThreshold,
		pin:          dataPin,
		voltage:      voltagePin,
	}
}

func (sensor *waterLevel) IsEmpty() bool {
	return sensor.adc.Get() <= sensor.dryThreshold
}

func (sensor *waterLevel) Configure() {
	sensor.adc = machine.ADC{Pin: sensor.pin}
	sensor.adc.Configure()

	sensor.voltage.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.voltage.Low()
}

func (sensor *waterLevel) On() {
	sensor.voltage.High()
}

func (sensor *waterLevel) Off() {
	sensor.voltage.Low()
}
