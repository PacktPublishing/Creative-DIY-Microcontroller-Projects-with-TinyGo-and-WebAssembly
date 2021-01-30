// Datasheet: https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf
package max7219spi

import (
	"machine"
)

type Driver interface {
	Configure()
	WriteCommand(register, data byte)

	StartShutdownMode()
	StopShutdownMode()
	StartDisplayTest()
	StopDisplayTest()
	SetDecodeMode(digitNumber uint8)
	SetScanLimit(digitNumber uint8)
}

type driver struct {
	bus  machine.SPI
	din  machine.Pin // din
	load machine.Pin // load
	clk  machine.Pin // clk
}

// NewDriver creates a new max7219 connection. The SPI wire must already be configured
func NewDriver(load machine.Pin, bus machine.SPI) Driver {
	return &driver{
		load: load,
		bus:  bus,
	}
}

func (driver *driver) Configure() {
	outPutConfig := machine.PinConfig{Mode: machine.PinOutput}

	driver.din.Configure(outPutConfig)
	driver.load.Configure(outPutConfig)
	driver.clk.Configure(outPutConfig)
}

func (driver *driver) SetScanLimit(digitNumber uint8) {
	driver.WriteCommand(byte(REG_SCANLIMIT), byte(digitNumber-1))
}

func (driver *driver) SetDecodeMode(digitNumber uint8) {
	switch digitNumber {
	case 1: // only decode first digit
		driver.WriteCommand(byte(REG_DECODE_MODE), 0x01)
	case 2, 3, 4: //  decode digits 3-0
		driver.WriteCommand(byte(REG_DECODE_MODE), 0x0F)
	case 8: // decode 8 digits
		driver.WriteCommand(byte(REG_DECODE_MODE), 0xFF)
	}
}

func (driver *driver) StartShutdownMode() {
	driver.WriteCommand(byte(REG_SHUTDOWN), 0x00)

}

func (driver *driver) StopShutdownMode() {
	driver.WriteCommand(byte(REG_SHUTDOWN), 0x01)
}

func (driver *driver) StartDisplayTest() {
	driver.WriteCommand(byte(REG_DISPLAY_TEST), 0x01)

}

func (driver *driver) StopDisplayTest() {
	driver.WriteCommand(byte(REG_DISPLAY_TEST), 0x00)
}

func (driver *driver) writeByte(data byte) {
	driver.bus.Transfer(data)
}

func (driver *driver) WriteCommand(register, data byte) {
	driver.load.Low()
	driver.writeByte(register)
	driver.writeByte(data)
	driver.load.High()
}
