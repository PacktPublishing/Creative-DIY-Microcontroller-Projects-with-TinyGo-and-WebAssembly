package main

import (
	"machine"
	"time"

	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter04/buzzer"
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter05/hs42561k"
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter05/max7219spi"
	hcsr04 "github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter05/ultrasonic-distance-sensor"
)

func main() {
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

	displayDriver := max7219spi.NewDevice(machine.D6, machine.SPI0)
	displayDriver.Configure()
	display := hs42561k.NewDevice(displayDriver, 4)
	display.Configure()

	distanceSensor := hcsr04.NewDevice(machine.D2, machine.D3, 60)
	distanceSensor.Configure()

	buzzer := buzzer.NewBuzzer(machine.D5)
	buzzer.Configure()

	println("all devices configured")

	for {
		currentDistance := distanceSensor.GetDistance()
		println("current distance:", currentDistance)
		if currentDistance >= 12 && currentDistance <= 25 {
			println("timer activated")
			handleTimer(display, displayDriver, buzzer)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func handleTimer(display hs42561k.Device, displayDriver max7219spi.Device, buzzer buzzer.Buzzer) {
	display.Configure()
	buzzer.Beep(100*time.Millisecond, 2)

	for i := 120; i > 0; i-- {
		println("counting:", i)
		if i >= 100 {
			display.SetDigit(2, hs42561k.Character(i/100))
		} else {
			display.SetDigit(2, hs42561k.Blank)
		}

		if i >= 10 {
			display.SetDigit(3, hs42561k.Character(i/10))
			if i%10 == 0 {
				display.SetDigit(4, hs42561k.Character(0))
			} else {
				display.SetDigit(4, hs42561k.Character(i-10))
			}
		} else {
			display.SetDigit(3, hs42561k.Blank)
			display.SetDigit(4, hs42561k.Character(i))
		}
		time.Sleep(time.Second)
	}

	display.SetDigit(3, hs42561k.Blank)
	display.SetDigit(4, hs42561k.Blank)

	buzzer.Beep(500*time.Millisecond, 1)
	displayDriver.StartShutdownMode()
}
