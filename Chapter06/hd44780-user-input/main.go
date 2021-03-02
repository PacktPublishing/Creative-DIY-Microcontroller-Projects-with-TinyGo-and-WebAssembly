package main

import (
	"machine"

	"github.com/Nerzal/drivers/hd44780i2c"
)

const carriageReturn = 0x0D

var (
	uart = machine.UART0
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(machine.I2C0, 0x27) // some modules have address 0x3F

	err := lcd.Configure(hd44780i2c.Config{
		Width:       16, // required
		Height:      2,  // required
		CursorOn:    false,
		CursorBlink: false,
	})
	if err != nil {
		println("failed to configure display")
	}

	lcd.Print([]byte(" Type to print "))

	hadInput := false
	for {
		if uart.Buffered() == 0 {
			continue
		}

		if !hadInput {
			hadInput = true
			lcd.ClearDisplay()
		}

		data, err := uart.ReadByte()
		if err != nil {
			println(err.Error())
		}

		if data == carriageReturn {
			lcd.Print([]byte("\n"))
			uart.Write([]byte("\n"))
			continue
		}

		lcd.Print([]byte{data})
		uart.WriteByte(data)
	}
}
