package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(machine.I2C0, 0x27) // some modules have address 0x3F

	lcd.Configure(hd44780i2c.Config{
		Width:       16, // required
		Height:      2,  // required
		CursorOn:    false,
		CursorBlink: false,
	})

	println("print first text")
	lcd.Print([]byte(" Hello World \n LCD 16x02"))
	time.Sleep(3 * time.Second)

}

func animation(lcd hd44780i2c.Device) {
	animation := []byte(" Hello World \n Send by \n Arduino Nano \n 33 IoT \n  powered by  \n TinyGo")
	lcd.ClearDisplay()
	for {
		for i := range animation {
			lcd.Print([]byte(string(animation[i])))
			time.Sleep(150 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)
		lcd.ClearDisplay()

	}
}
