// Datasheet: https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf
package max7219

import (
	"machine"
	"time"
)

type Driver interface {
	Configure()
	WriteCommand(register, data byte)
}

type driver struct {
	din  machine.Pin // din
	load machine.Pin // load
	clk  machine.Pin // clk
}

func NewDriver(din, load, clk machine.Pin) Driver {
	return &driver{
		din:  din,
		load: load,
		clk:  clk,
	}
}

func (driver *driver) Configure() {
	outPutConfig := machine.PinConfig{Mode: machine.PinOutput}

	driver.din.Configure(outPutConfig)
	driver.load.Configure(outPutConfig)
	driver.clk.Configure(outPutConfig)
}

func (driver *driver) writeByte(data byte) {
	for i := 8; i > 0; i-- {
		bitMask := byte(1 << (i - 1))
		driver.clk.Low()
		if (data & bitMask) != 0 {
			driver.din.High()
		} else {
			driver.din.Low()
		}
		driver.clk.High()
		time.Sleep(100 * time.Nanosecond)
	}
}

func (driver *driver) WriteCommand(register, data byte) {
	driver.load.Low()
	driver.writeByte(register)
	driver.writeByte(data)
	driver.load.High()
}
