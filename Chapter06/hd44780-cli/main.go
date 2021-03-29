package main

import (
	"machine"

	"tinygo.org/x/drivers/hd44780i2c"
)

const (
	carriageReturn = 0x0D
	homeCommand    = "#home"
	clearCommand   = "#clear"
)

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

	homeScreen(lcd)

	var commandBuffer string
	var commandIndex uint8

	commandStart := false
	hadInput := false
	for {
		if uart.Buffered() == 0 {
			continue
		}

		if !hadInput {
			hadInput = true
			clearDisplay(lcd)
		}

		data, err := uart.ReadByte()
		if err != nil {
			println(err.Error())
		}

		if string(data) == "#" {
			commandStart = true
			uart.Write([]byte("\ncommand started\n"))
		}

		if commandStart {
			commandBuffer += string(data)
			commandIndex++
		}

		switch commandBuffer {
		case homeCommand:
			uart.WriteByte(data)
			homeScreen(lcd)
			commandStart = false
			commandIndex = 0
			commandBuffer = ""
			continue
		case clearCommand:
			uart.WriteByte(data)
			clearDisplay(lcd)
			commandStart = false
			commandIndex = 0
			commandBuffer = ""
			continue
		}

		if commandIndex > 5 {
			commandStart = false
			commandIndex = 0
			commandBuffer = ""
			uart.Write([]byte("\nresetting command state\n"))
		}

		if data == carriageReturn {
			lcd.Print([]byte("\n"))
			uart.Write([]byte("\r\n"))
			continue
		}

		lcd.Print([]byte{data})
		uart.WriteByte(data)
	}
}

func homeScreen(lcd hd44780i2c.Device) {
	println("\nexecuting command homescreen\n")
	clearDisplay(lcd)
	lcd.Print([]byte(" TinyGo UART \n CLI "))
}

func clearDisplay(lcd hd44780i2c.Device) {
	println("\nexecuting command cleardisplay\n")
	lcd.ClearDisplay()
}
