package hcsr04

import (
	"machine"
	"time"
)

const speedOfSound = 0.034

type HCSR04 interface {
	Configure()
	GetDistance() uint16
	GetDistanceFromPulseLength(pulseLength float32) uint16
}

type hcsr04 struct {
	trigger machine.Pin
	echo    machine.Pin
	timeout uint16
}

func NewHCSR04(trigger, echo machine.Pin, maxDistance float32) HCSR04 {
	timeout := uint16(maxDistance / speedOfSound)

	return &hcsr04{
		trigger: trigger,
		echo:    echo,
		timeout: timeout,
	}
}

func (sensor *hcsr04) Configure() {
	sensor.trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.trigger.Configure(machine.PinConfig{Mode: machine.PinInput})
}

func (sensor *hcsr04) GetDistance() uint16 {
	timeoutTimer := time.Now()

	sensor.trigger.High()
	time.Sleep(10 * time.Microsecond)
	sensor.trigger.Low()
	for {
		if sensor.echo.Get() {
			timeoutTimer = time.Now()
			break
		}

		if time.Since(timeoutTimer).Microseconds() > int64(sensor.timeout) {
			return 0
		}
	}

	var pulseLength float32
	for {
		if !sensor.echo.Get() {
			pulseLength = float32(time.Since(timeoutTimer).Microseconds())
			break
		}

		if time.Since(timeoutTimer).Microseconds() > int64(sensor.timeout) {
			return 0
		}
	}

	return sensor.GetDistanceFromPulseLength(pulseLength)
}

func (sensor *hcsr04) GetDistanceFromPulseLength(pulseLength float32) uint16 {
	pulseLength = pulseLength / 2
	result := pulseLength * speedOfSound

	return uint16(result)
}
