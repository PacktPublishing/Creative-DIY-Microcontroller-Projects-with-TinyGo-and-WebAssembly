package pump

import (
	"machine"
	"time"
)

// Pump is used to control a pump.
type Pump interface {
	Configure()
	Pump(duration time.Duration, iterations uint8)
}

type pump struct {
	pin machine.Pin
}

// NewPump creates a new instance of Pump.
func NewPump(pin machine.Pin) Pump {
	return &pump{
		pin: pin,
	}
}

func (pump *pump) Configure() {
	pump.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.D5.Low()
}

func (pump *pump) Pump(duration time.Duration, iterations uint8) {
	for i := iterations; i > 0; i-- {
		pump.pin.High()
		time.Sleep(duration)
		pump.pin.Low()
		time.Sleep(duration)
	}
}
