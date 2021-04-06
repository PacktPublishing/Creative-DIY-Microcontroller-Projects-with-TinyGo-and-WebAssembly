package servo

import (
	"machine"
	"time"
)

// center
const centerDutyCycle = 1500 * time.Microsecond
const centerRemainingPeriod = 18500 * time.Microsecond

// left
const leftDutyCycle = 1000 * time.Microsecond
const leftRemainingPeriod = 19000 * time.Microsecond

// right
const rightDutyCycle = 2000 * time.Microsecond
const rightRemainingPeriod = 18000 * time.Microsecond

// Driver is used to control a sg90 servomotor.
type Driver struct {
	pin machine.Pin
}

func (servo *Driver) Configure(pin machine.Pin) {
	servo.pin = pin
	servo.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (servo *Driver) Right() {
	// prevent jamming, so only rotate a bit
	for position := 0; position <= 4; position++ {
		servo.pin.High()
		time.Sleep(rightDutyCycle)
		servo.pin.Low()
		time.Sleep(rightRemainingPeriod)
	}
}

func (servo *Driver) Center() {
	for position := 0; position < 90; position++ {
		servo.pin.High()
		time.Sleep(centerDutyCycle)
		servo.pin.Low()
		time.Sleep(centerRemainingPeriod)
	}
}

func (servo *Driver) Left() {
	for position := 0; position < 180; position++ {
		servo.pin.High()
		time.Sleep(centerDutyCycle)
		servo.pin.Low()
		time.Sleep(centerRemainingPeriod)
	}
}
