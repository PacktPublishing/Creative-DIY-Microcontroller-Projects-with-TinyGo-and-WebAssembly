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

// Servo is used to control a sg90 servomotor.
type Servo struct {
	pwm machine.PWM
}

func (servo *Servo) Configure(pin machine.Pin) {
	machine.InitPWM()

	servoPin := machine.PWM{Pin: pin}
	servoPin.Configure()

	servo.pwm = servoPin
}

func (servo *Servo) Right() {
	// prevent jamming, so only rotate a bit
	for position := 0; position <= 4; position++ {
		servo.pwm.Pin.High()
		time.Sleep(rightDutyCycle)
		servo.pwm.Pin.Low()
		time.Sleep(rightRemainingPeriod)
	}
}

func (servo *Servo) Center() {
	for position := 0; position < 90; position++ {
		servo.pwm.Pin.High()
		time.Sleep(centerDutyCycle)
		servo.pwm.Pin.Low()
		time.Sleep(centerRemainingPeriod)
	}
}

func (servo *Servo) Left() {
	for position := 0; position < 180; position++ {
		servo.pwm.Pin.High()
		time.Sleep(centerDutyCycle)
		servo.pwm.Pin.Low()
		time.Sleep(centerRemainingPeriod)
	}
}
