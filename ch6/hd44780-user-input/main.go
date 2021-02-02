package main

import (
	"machine"
	"time"

	"github.com/Nerzal/drivers/hd44780i2c"
)

var (
	uart = machine.UART0
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
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

	time.Sleep(5 * time.Second)
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})

	println("print first text")
	lcd.Print([]byte(" Type to print "))
	time.Sleep(3 * time.Second)
	lcd.ClearDisplay()

	uart.Write([]byte("Echo console enabled. Type something then press enter:\r\n"))

	i := 1
	var input []byte
	for {
		uart.Read(input)
		if uart.Buffered() > 0 {
			data, err := uart.ReadByte()
			if err != nil {
				println(err.Error())
			}

			lcd.Print([]byte{data})
			uart.WriteByte(data)
			i++
		}

		if i == 16 {
			lcd.Print([]byte("\n"))
			uart.Write([]byte("\n"))
		}

		if i >= 32 {
			i = 0
			lcd.ClearDisplay()
		}
	}
}
