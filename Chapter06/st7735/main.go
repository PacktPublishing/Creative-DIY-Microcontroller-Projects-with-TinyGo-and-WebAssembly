package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"
)

var (
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
	black = color.RGBA{0, 0, 0, 255}
)

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
	})

	resetPin := machine.D6
	dcPin := machine.D5
	csPin := machine.D7
	backLightPin := machine.D2

	display := st7735.New(machine.SPI0, resetPin, dcPin, csPin, backLightPin)
	display.Configure(st7735.Config{})

	width, height := display.Size()

	display.FillRectangle(0, 0, width/2, height/2, white)
	display.FillRectangle(width/2, 0, width/2, height/2, red)
	display.FillRectangle(0, height/2, width/2, height/2, green)
	display.FillRectangle(width/2, height/2, width/2, height/2, blue)

	time.Sleep(3 * time.Second)
	// display.FillScreen(black)

	// i := 0
	// for {
	// 	display.FillScreen(black)

	// 	tinydraw.FilledRectangle(&display, 0, 0, 128, 32, white)

	// 	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 110, 145, "TinyDraw", red, tinyfont.ROTATION_180)

	// 	tinydraw.FilledCircle(&display, 64, 96, 32, green)
	// 	tinydraw.FilledCircle(&display, 64, 96, 24, blue)
	// 	tinydraw.FilledCircle(&display, 64, 96, 16, red)

	// 	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 110, 40, "TinyFont", green, tinyfont.ROTATION_180)

	// 	counterText := fmt.Sprintf("Count: %v", i)
	// 	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 123, 2, counterText, black, tinyfont.ROTATION_180)

	// 	time.Sleep(2 * time.Second)
	// 	i++
	// }
}
