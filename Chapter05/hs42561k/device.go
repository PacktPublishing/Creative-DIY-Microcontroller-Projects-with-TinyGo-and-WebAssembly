package hs42561k

import (
	"errors"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter05/max7219spi"
)

type Device interface {
	Configure()
	SetDigit(digit byte, character Character) error
}

type device struct {
	digitNumber   uint8
	displayDevice max7219spi.Device
}

func NewDevice(displayDevice max7219spi.Device, digitNumber uint8) Device {
	return &device{
		displayDevice: displayDevice,
		digitNumber:   digitNumber,
	}
}

func (device *device) Configure() {
	device.displayDevice.StopDisplayTest()
	device.displayDevice.SetDecodeMode(device.digitNumber)
	device.displayDevice.SetScanLimit(device.digitNumber)
	device.displayDevice.StopShutdownMode()

	for i := 1; i < int(device.digitNumber); i++ {
		device.displayDevice.WriteCommand(byte(i), byte(Blank))
	}
}

var ErrIllegalDigit = errors.New("Invalid digit selected")

func (device *device) SetDigit(digit byte, character Character) error {
	if uint8(digit) > device.digitNumber {
		return ErrIllegalDigit
	}

	device.displayDevice.WriteCommand(digit, byte(character))
	return nil
}
