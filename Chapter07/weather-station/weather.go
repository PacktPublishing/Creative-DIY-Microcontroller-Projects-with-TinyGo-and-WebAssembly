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

var (
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

type Service interface {
	CheckSensorConnectivity()
	ReadData() (temperature, pressure, humidity int32, err error)
	DisplayData(temperature, pressure, humidity int32)
	GetFormattedReadings(temperature, pressure, humidity int32) (temp, press, hum string)
	SavePressureReading(pressure float64)
	CheckAlert(alertThreshold float64, timeSpan int8) (bool, float64)
}

type service struct {
	sensor            *bme280.Device
	display           *st7735.Device
	readings          [6]float64
	readingsIndex     int8
	firstReadingSaved bool
}

func New(sensor *bme280.Device, display *st7735.Device) Service {
	return &service{
		sensor:            sensor,
		display:           display,
		readingsIndex:     int8(0),
		readings:          [6]float64{},
		firstReadingSaved: false,
	}
}

func (service *service) ReadData() (temp, press, hum int32, err error) {
	temp, err = service.sensor.ReadTemperature()
	if err != nil {
		return
	}

	press, err = service.sensor.ReadPressure()
	if err != nil {
		return
	}

	hum, err = service.sensor.ReadHumidity()
	if err != nil {
		return
	}

	return
}

func (service *service) CheckSensorConnectivity() {
	for {
		connected := service.sensor.Connected()
		if !connected {
			println("could not detect BME280")
			time.Sleep(time.Second)
		}

		println("BME280 detected")
		break
	}
}

func (service *service) DisplayData(temperature, pressure, humidity int32) {
	println("fill screen")
	service.display.FillScreen(black)

	tinyfont.WriteLineRotated(service.display, &freemono.Bold9pt7b, 110, 3, "Tiny Weather", white, tinyfont.ROTATION_90)

	temp, press, hum := service.GetFormattedReadings(temperature, pressure, humidity)

	tempString := "Temp:" + temp + "C"
	tinyfont.WriteLineRotated(service.display, &freemono.Bold9pt7b, 65, 3, tempString, white, tinyfont.ROTATION_90)

	pressString := "P:" + press + "hPa"
	tinyfont.WriteLineRotated(service.display, &freemono.Bold9pt7b, 45, 3, pressString, white, tinyfont.ROTATION_90)

	humString := "Hum:" + hum + "%"
	tinyfont.WriteLineRotated(service.display, &freemono.Bold9pt7b, 25, 3, humString, white, tinyfont.ROTATION_90)
}

func (service *service) GetFormattedReadings(temperature, pressure, humidity int32) (temp, press, hum string) {
	temp = strconv.FormatFloat(float64(temperature/1000), 'f', 2, 64)
	press = strconv.FormatFloat(float64(pressure/100000), 'f', 2, 64)
	hum = strconv.FormatFloat(float64(humidity/100), 'f', 2, 64)
	return
}

func (service *service) SavePressureReading(pressure float64) {
	if !service.firstReadingSaved {
		for i := 0; i < len(service.readings); i++ {
			service.readings[i] = pressure
		}

		service.firstReadingSaved = true
		service.readingsIndex = 0

		return
	}

	service.readingsIndex++
	service.readingsIndex = service.readingsIndex % int8(len(service.readings))
	service.readings[service.readingsIndex] = pressure
}

func (service *service) CheckAlert(alertThreshold float64, timeSpan int8) (bool, float64) {
	currentReading := service.readings[service.readingsIndex]

	comparisonIndex := service.readingsIndex - timeSpan
	if comparisonIndex < 0 {
		comparisonIndex = int8(len(service.readings)) - 1
	}

	println("comparing index:", service.readingsIndex, "with index: ", comparisonIndex)

	comparisonReading := service.readings[comparisonIndex]

	diff := comparisonReading - currentReading

	return diff >= alertThreshold, diff
}
