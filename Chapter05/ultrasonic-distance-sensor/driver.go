package hcsr04

import (
	"machine"
	"time"
)

const speedOfSound = 0.0343 // cm / us

type Device interface {
	Configure()
	GetDistance() uint16
	GetDistanceFromPulseLength(pulseLength float32) uint16
}

type device struct {
	trigger machine.Pin
	echo    machine.Pin
	timeout int64
}

func NewDevice(trigger, echo machine.Pin, maxDistance float32) Device {
	timeout := int64(maxDistance * 2 / speedOfSound)

	return &device{
		trigger: trigger,
		echo:    echo,
		timeout: timeout,
	}
}

func (sensor *device) Configure() {
	sensor.trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
}

func (sensor *device) sendPulse() {
	sensor.trigger.High()
	time.Sleep(10 * time.Microsecond)
	sensor.trigger.Low()
}

func (sensor *device) GetDistance() uint16 {
	i := 0
	timeoutTimer := time.Now()
	sensor.sendPulse()

	for {
		if sensor.echo.Get() {
			timeoutTimer = time.Now()
			break
		}
		i++
		if i > 15 {
			microseconds := time.Since(timeoutTimer).Microseconds()
			if microseconds > sensor.timeout {
				return 0
			}
		}
	}

	var pulseLength float32
	i = 0
	for {
		if !sensor.echo.Get() {
			microseconds := time.Since(timeoutTimer).Microseconds()
			pulseLength = float32(microseconds)
			break
		}

		i++
		if i > 15 {
			microseconds := time.Since(timeoutTimer).Microseconds()
			if microseconds > sensor.timeout {
				return 0
			}
		}
	}

	return sensor.GetDistanceFromPulseLength(pulseLength)
}

func (sensor *device) GetDistanceFromPulseLength(pulseLength float32) uint16 {
	pulseLength = pulseLength / 2
	result := pulseLength * speedOfSound

	return uint16(result)
}
