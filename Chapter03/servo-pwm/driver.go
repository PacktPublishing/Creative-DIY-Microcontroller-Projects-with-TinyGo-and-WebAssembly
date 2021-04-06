package servopwm

import (
	"machine"
)

const period = 20e6

type Device struct {
	pwm     machine.PWM
	pin     machine.Pin
	channel uint8
}

func NewDevice(timer machine.PWM, pin machine.Pin) *Device {
	return &Device{
		pwm: timer,
		pin: pin,
	}
}

func (d *Device) Configure() error {
	err := d.pwm.Configure(machine.PWMConfig{
		Period: period,
	})
	if err != nil {
		return err
	}

	d.channel, err = d.pwm.Channel(machine.Pin(d.pin))
	if err != nil {
		return err
	}

	return nil
}

func (d *Device) Right() {
	d.setDutyCycle(1000)
}

func (d *Device) Center() {
	d.setDutyCycle(1500)
}

func (d *Device) Left() {
	d.setDutyCycle(2000)
}

// setDutyCycle sets the cycle in microseconds
func (d *Device) setDutyCycle(cycle uint64) {
	value := uint64(d.pwm.Top()) * cycle / (period / 1000)
	d.pwm.Set(d.channel, uint32(value))
}
