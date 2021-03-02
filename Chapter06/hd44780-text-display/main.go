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
		Width:  16, // required
		Height: 2,  // required
	})

	lcd.Print([]byte(" Hello World \n LCD 16x02"))

	time.Sleep(5 * time.Second)
	lcd.Print([]byte("We just print more text, to see what happens, when we overflow the 16x2 character limit"))

	time.Sleep(5 * time.Second)
	animation(lcd)

}

func animation(lcd hd44780i2c.Device) {
	text := []byte(" Hello World \n Send by \n Arduino Nano \n 33 IoT \n  powered by  \n TinyGo")
	lcd.ClearDisplay()
	for {
		for i := range text {
			lcd.Print([]byte(string(text[i])))
			time.Sleep(150 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)
		lcd.ClearDisplay()
	}
}
