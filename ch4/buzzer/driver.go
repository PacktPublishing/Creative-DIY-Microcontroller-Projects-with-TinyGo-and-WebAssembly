package buzzer

import (
	"machine"
	"time"
)

// Buzzer is used to control a buzzer
type Buzzer interface {
	Configure()
	Beep(highDuration time.Duration, amount uint8)
}

type buzzer struct {
	pin machine.Pin
}

// NewBuzzer returns a new instance of buzzer
func NewBuzzer(pin machine.Pin) Buzzer {
	return buzzer{pin: pin}
}

// Configure Buzzer as outout
func (buzzer buzzer) Configure() {
	buzzer.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (buzzer buzzer) Beep(highDuration time.Duration, amount uint8) {
	for i := amount; i > 0; i-- {
		buzzer.pin.High()
		time.Sleep(highDuration)
		buzzer.pin.Low()
		time.Sleep(highDuration)
	}
}
