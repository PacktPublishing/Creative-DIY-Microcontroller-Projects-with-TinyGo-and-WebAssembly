package weatherstation_test

import (
	"testing"

	weatherstation "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/weather-station"
)

func Test_Alert_NoAlerts(t *testing.T) {
	// Arrange
	service := weatherstation.New(nil, nil)
	service.SavePressureReading(1000.54)
	service.SavePressureReading(1000.54)
	service.SavePressureReading(1000.54)
	service.SavePressureReading(1000.54)
	service.SavePressureReading(1000.54)
	service.SavePressureReading(1000.54)

	// Act
	t.Run("no alert, 1 hour timespan", func(t *testing.T) {
		alert, _ := service.CheckAlert(2, 1)
		if alert { // Assert
			t.Error("calculated an alert, but there should be none")
		}
	})

	// Act
	t.Run("no alert, 3 hour timespan", func(t *testing.T) {
		alert, _ := service.CheckAlert(6, 3)
		if alert { // Assert
			t.Error("calculated an alert, but there should be none")
		}
	})
}

func Test_Alert(t *testing.T) {
	// Arrange
	service := weatherstation.New(nil, nil)
	service.SavePressureReading(1000.54)
	service.SavePressureReading(950.23)

	// Act
	t.Run("alert, 1 hour timespan", func(t *testing.T) {
		alert, diff := service.CheckAlert(2, 1)
		if !alert { // Assert
			t.Errorf("calculated no alert, but there should be one. diff was: %f", diff)
		} else {
			t.Logf("calculated alert for diff: %f", diff)

		}

	})

	service.SavePressureReading(950.23)
	service.SavePressureReading(950.23)

	// Act
	t.Run("no alert, 3 hour timespan", func(t *testing.T) {
		alert, diff := service.CheckAlert(6, 3)
		if !alert { // Assert
			t.Errorf("calculated no alert, but there should be one. diff was: %f", diff)
		} else {
			t.Logf("calculated alert for diff: %f", diff)

		}

	})
}
