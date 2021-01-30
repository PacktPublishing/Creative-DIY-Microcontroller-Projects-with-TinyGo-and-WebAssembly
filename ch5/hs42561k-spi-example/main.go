package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch5/hs42561k"
	max7219spi "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch5/max7219-spi"
)

var characters = []hs42561k.Character{
	hs42561k.Zero,
	hs42561k.One,
	hs42561k.Two,
	hs42561k.Three,
	hs42561k.Four,
	hs42561k.Five,
	hs42561k.Six,
	hs42561k.Seven,
	hs42561k.Eight,
	hs42561k.Nine,
	hs42561k.Dash,
	hs42561k.E,
	hs42561k.H,
	hs42561k.L,
	hs42561k.P,
	hs42561k.Blank,
	hs42561k.Dot,
}

func main() {
	time.Sleep(2 * time.Second)
	println("startup")

	err := machine.SPI0.Configure(machine.SPIConfig{
		SDO:       machine.D11,
		SCK:       machine.D13,
		LSBFirst:  false,
		Frequency: 10000000,
	})

	if err != nil {
		println("failed to configure spi:", err.Error())
	}

	println("spi configured")

	displayDriver := max7219spi.NewDriver(machine.D6, machine.SPI0)
	displayDriver.Configure()
	display := hs42561k.NewDriver(displayDriver, 4)
	display.Configure()

	println("display configured")

	for {
		for _, character := range characters {
			println("writing", "characterValue:", character.String())
			display.SetDigit(4, character)
			display.SetDigit(3, character)
			display.SetDigit(2, character)
			display.SetDigit(1, character)
			time.Sleep(500 * time.Millisecond)

		}
		time.Sleep(time.Second)
	}
}
