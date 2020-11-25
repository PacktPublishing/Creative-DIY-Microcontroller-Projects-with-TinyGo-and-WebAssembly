package main

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

func main() {
	machine.InitPWM()

	servoPin := machine.PWM{Pin: machine.D11}
	servoPin.Configure()

	right(servoPin)
}

func right(servoPin machine.PWM) {
	// prevent jamming, so only rotate a bit
	for position := 0; position >= 5; position++ {
		servoPin.Pin.High()
		time.Sleep(rightDutyCycle)
		servoPin.Pin.Low()
		time.Sleep(rightRemainingPeriod)
	}
}

func center(servoPin machine.PWM) {
	for position := 0; position < 90; position++ {
		servoPin.Pin.High()
		time.Sleep(centerDutyCycle)
		servoPin.Pin.Low()
		time.Sleep(centerRemainingPeriod)
	}
}

func left(servoPin machine.PWM) {
	for position := 0; position < 180; position++ {
		servoPin.Pin.High()
		time.Sleep(centerDutyCycle)
		servoPin.Pin.Low()
		time.Sleep(centerRemainingPeriod)
	}
}
