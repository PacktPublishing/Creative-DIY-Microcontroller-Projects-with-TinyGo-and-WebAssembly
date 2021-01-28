package main

import (
	"machine"
	"time"

	hcsr04 "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch5/ultrasonic-distance-sensor"
)

func main() {
	sensor := hcsr04.NewHCSR04(machine.D2, machine.D3, 400)
	sensor.Configure()

	for {
		distance := sensor.GetDistance()
		if distance != 0 {
			println("Current distance: ", distance, "cm")
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
