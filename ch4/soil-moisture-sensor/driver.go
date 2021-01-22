package soil

import (
	"machine"
)

// SoilSensor is used to read the humidity value of a soil sensor
type SoilSensor interface {
	Get() MoistureLevel
	Configure()
	On()
	Off()
}

type soilSensor struct {
	waterThreshold         float32
	completelyDryThreshold float32
	category               float32
	pin                    machine.Pin
	adc                    machine.ADC
	voltage                machine.Pin
}

type MoistureLevel uint8

const (
	CompletelyDry MoistureLevel = iota
	VeryDry       MoistureLevel = iota
	Dry           MoistureLevel = iota
	Wet           MoistureLevel = iota
	VeryWet       MoistureLevel = iota
	Water         MoistureLevel = iota
)

func NewSoilSensor(waterThreshold, dryThreshold float32, dataPin, voltagePin machine.Pin) SoilSensor {
	category := (dryThreshold - waterThreshold) / 6
	return &soilSensor{
		waterThreshold:         waterThreshold,
		completelyDryThreshold: dryThreshold,
		category:               category,
		pin:                    dataPin,
		voltage:                voltagePin,
	}
}

func (sensor *soilSensor) Get() MoistureLevel {
	value := float32(sensor.adc.Get())

	switch {
	case value >= sensor.completelyDryThreshold:
		return CompletelyDry
	case value <= sensor.completelyDryThreshold-sensor.category &&
		value > sensor.completelyDryThreshold-sensor.category*2:
		return VeryDry
	case value <= sensor.completelyDryThreshold-sensor.category*2 &&
		value > sensor.completelyDryThreshold-sensor.category*3:
		return Dry
	case value <= sensor.completelyDryThreshold-sensor.category*4 &&
		value > sensor.completelyDryThreshold-sensor.category*5:
		return Wet
	case value <= sensor.completelyDryThreshold-sensor.category*5 &&
		value > sensor.completelyDryThreshold-sensor.category*6:
		return VeryWet
	default:
		return Water
	}
}

func (sensor *soilSensor) Configure() {
	sensor.adc = machine.ADC{Pin: sensor.pin}
	sensor.adc.Configure()

	sensor.voltage.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.voltage.Low()
}

func (sensor *soilSensor) On() {
	sensor.voltage.High()
}

func (sensor *soilSensor) Off() {
	sensor.voltage.Low()
}
