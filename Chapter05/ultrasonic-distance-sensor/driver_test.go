package hcsr04_test

import (
	"machine"
	"testing"

	hcsr04 "github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter05/ultrasonic-distance-sensor"
)

func TestGetDistanceFromPulseLength_30cm(t *testing.T) {
	sensor := hcsr04.NewDevice(machine.D2, machine.D3, 100)

	distance := sensor.GetDistanceFromPulseLength(1749.27113703)

	if distance != 30 {
		t.Error("Expected distance: 30cm", "actual distance: ", distance, "cm")
	}
}

func TestGetDistanceFromPulseLength_TableDriven(t *testing.T) {
	// Arrange
	var testCases = [4]struct {
		Name        string
		Result      uint16
		PulseLength float32
	}{
		{
			Name:        "1cm",
			Result:      1,
			PulseLength: 58.8235294117,
		},
		{
			Name:        "30cm",
			Result:      30,
			PulseLength: 1749.27113703,
		},
		{
			Name:        "60cm",
			Result:      60,
			PulseLength: 3498.54227405,
		},
		{
			Name:        "400cm",
			Result:      400,
			PulseLength: 23323.6151603,
		},
	}

	sensor := hcsr04.NewDevice(machine.D2, machine.D3, 100)

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			// Act
			distance := sensor.GetDistanceFromPulseLength(testCase.PulseLength)

			// Assert
			if distance != testCase.Result {
				t.Error("Expected distance:", testCase.Name, "actual distance: ", distance, "cm")
			}
		})
	}
}
