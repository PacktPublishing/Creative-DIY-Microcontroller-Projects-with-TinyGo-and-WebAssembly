package soil

import (
	"machine"
)

// SoilSensor is used to read the humidity value of a soil sensor
type SoilSensor interface {
	Get() uint8
	Configure()
	On()
	Off()
}

type soilSensor struct {
	wetThreshold float32
	dryThreshold float32
	pin          machine.Pin
	adc          machine.ADC
	voltage      machine.Pin
}

func NewSoilSensor(wetThreshold, dryThreshold float32, dataPin, voltagePin machine.Pin) SoilSensor {
	return soilSensor{
		wetThreshold: wetThreshold,
		dryThreshold: dryThreshold,
		pin:          dataPin,
		voltage:      voltagePin,
	}
}

func (sensor soilSensor) Get() uint8 {
	value := sensor.adc.Get()
	println("value reading:", value)

	if float32(value) <= sensor.wetThreshold {
		return 100
	}

	if float32(value) > sensor.dryThreshold {
		return 0
	}

	return uint8(float32(value) / sensor.dryThreshold * 100)
}

func (sensor soilSensor) Configure() {
	sensor.adc = machine.ADC{Pin: sensor.pin}
	sensor.adc.Configure()

	sensor.voltage.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.voltage.Low()
}

func (sensor soilSensor) On() {
	sensor.voltage.High()
}

func (sensor soilSensor) Off() {
	sensor.voltage.Low()
}
