// Datasheet: https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf
package max7219

import "machine"

type Driver interface {
}

type driver struct {
	din machine.Pin
	cs  machine.Pin
	clk machine.Pin
}

func NewDriver(din, cs, clk machine.Pin) Driver {
	return &driver{
		din: din,
		cs:  cs,
		clk: clk,
	}
}

func (driver *driver) Configure() {
	outPutConfig := machine.PinConfig{Mode: machine.PinOutput}

	driver.din.Configure(outPutConfig)
	driver.cs.Configure(outPutConfig)
	driver.clk.Configure(outPutConfig)
}
