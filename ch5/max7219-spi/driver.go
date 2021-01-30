// Datasheet: https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf
package max7219spi

import (
	"machine"
)

type Driver interface {
	Configure()
	WriteCommand(register, data byte)
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

func (driver *driver) writeByte(data byte) {
	driver.bus.Transfer(data)
}

func (driver *driver) WriteCommand(register, data byte) {
	driver.load.Low()
	driver.writeByte(register)
	driver.writeByte(data)
	driver.load.High()
}
