package hcsr04_test

import (
	"machine"
	"testing"

	hcsr04 "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch5/supersonic-distance-sensor"
)

func TestGetDistanceFromPulseLength_30cm(t *testing.T) {
	sensor := hcsr04.NewHCSR04(machine.D2, machine.D3, 100)

	distance := sensor.GetDistanceFromPulseLength(1764.70588235)

	if distance != 30 {
		t.Error("Expected distance: 30cm", "actual distance: ", distance, "cm")
	}
}

func TestGetDistanceFromPulseLength_TableDriven(t *testing.T) {
	var testCases = []struct {
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
			PulseLength: 1764.70588235,
		},
		{
			Name:        "60cm",
			Result:      60,
			PulseLength: 3529.4117647,
		},
		{
			Name:        "400cm",
			Result:      400,
			PulseLength: 23529.4117647,
		},
	}

	sensor := hcsr04.NewHCSR04(machine.D2, machine.D3, 100)

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			distance := sensor.GetDistanceFromPulseLength(testCase.PulseLength)

			if distance != testCase.Result {
				t.Error("Expected distance: 30cm", "actual distance: ", distance, "cm")
			}
		})
	}
}
