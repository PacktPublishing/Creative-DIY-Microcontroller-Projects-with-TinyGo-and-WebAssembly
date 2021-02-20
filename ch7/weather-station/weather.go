package weatherstation

import (
	"image/color"
	"strconv"
	"time"

	"tinygo.org/x/drivers/bme280"
	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

// pressure loss of more than 6 hPa / 3h -> possible storm incoming
// 							  2 hpa / 1h -> possible storm incoming

var (
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

type WeatherStation interface {
	CheckSensorConnectivity()
	ReadData() (temperature, pressure, humidity, altitude int32, err error)
	DisplayData(temperature, pressure, humidity, altitude int32)
}

type weatherStation struct {
	sensor          bme280.Device
	display         st7735.Device
	monitorReadings bool
}

func New(sensor bme280.Device, display st7735.Device) WeatherStation {
	return &weatherStation{
		sensor:  sensor,
		display: display,
	}
}

func (weatherStation *weatherStation) ReadData() (temperature, pressure, humidity, altitude int32, err error) {
	temperature, err = weatherStation.sensor.ReadTemperature()
	if err != nil {
		return
	}

	pressure, err = weatherStation.sensor.ReadPressure()
	if err != nil {
		return
	}

	humidity, err = weatherStation.sensor.ReadHumidity()
	if err != nil {
		return
	}

	altitude, err = weatherStation.sensor.ReadAltitude()
	if err != nil {
		return
	}

	return
}

func (weatherStation *weatherStation) CheckSensorConnectivity() {
	for {
		connected := weatherStation.sensor.Connected()
		if !connected {
			println("could not BME280 not detect")
			time.Sleep(time.Second)
		}

		println("BME280 detected")
		break
	}
}

func (weatherStation *weatherStation) DisplayData(temperature, pressure, humidity, altitude int32) {
	weatherStation.display.FillScreen(black)

	tinyfont.WriteLineRotated(&weatherStation.display, &freemono.Bold9pt7b, 110, 3, "Tiny Weather", white, tinyfont.ROTATION_90)

	tempString := "Temp:" + strconv.FormatFloat(float64(temperature)/1000, 'f', 2, 64) + "C"
	tinyfont.WriteLineRotated(&weatherStation.display, &freemono.Bold9pt7b, 65, 3, tempString, white, tinyfont.ROTATION_90)

	pressString := "P:" + strconv.FormatFloat(float64(pressure)/100000, 'f', 2, 64) + "hPa"
	tinyfont.WriteLineRotated(&weatherStation.display, &freemono.Bold9pt7b, 45, 3, pressString, white, tinyfont.ROTATION_90)

	humString := "Hum:" + strconv.FormatFloat(float64(humidity)/100, 'f', 2, 64) + "%"
	tinyfont.WriteLineRotated(&weatherStation.display, &freemono.Bold9pt7b, 25, 3, humString, white, tinyfont.ROTATION_90)

	altString := "Alt:" + strconv.FormatInt(int64(altitude), 10) + "m"
	tinyfont.WriteLineRotated(&weatherStation.display, &freemono.Bold9pt7b, 5, 3, altString, white, tinyfont.ROTATION_90)
}
