package main

import (
	"machine"
	"time"

	hcsr04 "github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter05/ultrasonic-distance-sensor"
)

func main() {
	sensor := hcsr04.NewDevice(machine.D2, machine.D3, 400)
	sensor.Configure()

	for {
		distance := sensor.GetDistance()
		if distance != 0 {
			println("Current distance: ", distance, "cm")
		} else {
			println("out of range")
		}

		time.Sleep(time.Second)
	}
}
