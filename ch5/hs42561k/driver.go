package hs42561k

import (
	"errors"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch5/max7219spi"
)

type Driver interface {
	Configure()
	SetDigit(digit byte, character Character) error
}

type driver struct {
	digitNumber   uint8
	displayDriver max7219spi.Device
}

func NewDriver(displayDriver max7219spi.Device, digitNumber uint8) Driver {
	return &driver{
		displayDriver: displayDriver,
		digitNumber:   digitNumber,
	}
}

func (driver *driver) Configure() {
	driver.displayDriver.StopDisplayTest()
	driver.displayDriver.SetDecodeMode(driver.digitNumber)
	driver.displayDriver.StopShutdownMode()
	driver.displayDriver.SetScanLimit(driver.digitNumber)

	for i := 1; i < int(driver.digitNumber); i++ {
		driver.displayDriver.WriteCommand(byte(i), byte(Blank))
	}
}

var ErrIllegalDigit = errors.New("Invalid digit selected")

func (driver *driver) SetDigit(digit byte, character Character) error {
	if uint8(digit) > driver.digitNumber {
		return ErrIllegalDigit
	}

	driver.displayDriver.WriteCommand(digit, byte(character))

	return nil
}
