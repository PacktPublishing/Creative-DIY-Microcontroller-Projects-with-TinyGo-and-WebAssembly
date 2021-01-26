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
	timeout int64
}

func NewHCSR04(trigger, echo machine.Pin, maxDistance float32) HCSR04 {
	timeout := int64(maxDistance/speedOfSound) * 10

	return &hcsr04{
		trigger: trigger,
		echo:    echo,
		timeout: timeout,
	}
}

func (sensor *hcsr04) Configure() {
	sensor.trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
}

func (sensor *hcsr04) GetDistance() uint16 {
	timeoutTimer := time.Now()

	sensor.trigger.Low()
	time.Sleep(2 * time.Microsecond)
	sensor.trigger.High()
	time.Sleep(10 * time.Microsecond)
	sensor.trigger.Low()
	for {
		if sensor.echo.Get() {
			timeoutTimer = time.Now()
			break
		}

		microseconds := time.Since(timeoutTimer).Microseconds()
		if microseconds > int64(sensor.timeout) {
			println("timeout:", sensor.timeout, "after:", microseconds)
			return 0
		}
	}

	var pulseLength float32
	i := 0
	for {
		if !sensor.echo.Get() {
			pulseLength = float32(time.Since(timeoutTimer).Microseconds() - 12)
			break
		}

		i++
		if i > 10 {
			microseconds := time.Since(timeoutTimer).Microseconds()
			if microseconds > int64(sensor.timeout) {
				println("timeout:", sensor.timeout, "after:", microseconds)
				return 0
			}
		}
	}

	return sensor.GetDistanceFromPulseLength(pulseLength)
}

func (sensor *hcsr04) GetDistanceFromPulseLength(pulseLength float32) uint16 {
	pulseLength = pulseLength / 2
	result := pulseLength * speedOfSound

	return uint16(result)
}
