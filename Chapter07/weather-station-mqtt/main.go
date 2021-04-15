package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/Nerzal/drivers/wifinina"
	mqttclient "github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter07/mqtt-client"
	weatherstation "github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter07/weather-station"
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter07/wifi"
	"tinygo.org/x/drivers/bme280"
)

const ssid = ""
const password = ""

var (
	temperature int32
	pressure    int32
	humidity    int32
)

func printError(message string, err error) {
	for {
		println(message, err.Error())
		time.Sleep(time.Second)
	}
}

func main() {
	time.Sleep(5 * time.Second)

	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := bme280.New(machine.I2C0)
	sensor.Configure()

	weatherStation := weatherstation.New(&sensor, nil)
	weatherStation.CheckSensorConnectivity()

	wifiClient := wifi.New(ssid, password)

	println("configuring nina wifi chip")
	err := wifiClient.Configure()
	if err != nil {
		printError("could not configure wifi client", err)
	}

	println("checking firmware")
	err = wifiClient.CheckHardware()
	if err != nil {
		printError("could not check hardware", err)
	}

	wifiClient.ConnectWifi()

	mqttClient := mqttclient.New("tcp://192.168.2.102:1883", "weatherStation")
	println("connecting to mqtt broker")
	err = mqttClient.ConnectBroker()
	if err != nil {
		printError("could not configure mqtt", err)
	}
	println("connected to mqtt broker")

	go publishSensorData(mqttClient, wifiClient, weatherStation)
	go publishAlert(mqttClient, wifiClient, weatherStation)

	for {
		temperature, pressure, humidity, err = weatherStation.ReadData()
		if err != nil {
			printError("could not read sensor data:", err)
		}

		time.Sleep(time.Minute)
	}

}

func publishSensorData(mqttClient *mqttclient.Client, wifiClient wifi.Client, weatherStation weatherstation.Service) {
	for {
		time.Sleep(time.Minute)
		println("publishing sensor data")

		tempString, pressString, humidityString := weatherStation.GetFormattedReadings(temperature, pressure, humidity)
		message := []byte(fmt.Sprintf("sensor readings#%s#%s#%s", tempString, pressString, humidityString))

		err := mqttClient.PublishMessage("weather/data", message, 0, true)
		if err != nil {
			println(err.Error())
			switch err.(type) {
			case wifinina.Error:
				wifiClient.ConnectWifi()
				mqttClient.ConnectBroker()
			}
		}
	}
}

func publishAlert(mqttClient *mqttclient.Client, wifiClient wifi.Client, weatherStation weatherstation.Service) {
	// pressure loss of more than 6 hPa / 3h -> possible storm incoming
	// 							  2 hpa / 1h -> possible storm incoming
	// source http://www.bohlken.net/airpressure2.htm
	for {
		time.Sleep(time.Hour)

		weatherStation.SavePressureReading(float64(pressure / 100000))

		alert, diff := weatherStation.CheckAlert(2, 1)
		if alert {
			err := mqttClient.PublishMessage("weather/alert", []byte(fmt.Sprintf("%s#%v#%s", "possible storm incoming", diff, "1 hour")), 0, true)
			if err != nil {
				switch err.(type) {
				case wifinina.Error:
					println(err.Error())
					wifiClient.ConnectWifi()
					mqttClient.ConnectBroker()
				default:
					println(err.Error())
				}
			}
		}

		alert, diff = weatherStation.CheckAlert(6, 3)
		if !alert {
			continue
		}

		err := mqttClient.PublishMessage("weather/alert", []byte(fmt.Sprintf("%s#%v#%s", "possible storm incoming", diff, "3 hours")), 0, true)
		if err != nil {
			switch err.(type) {
			case wifinina.Error:
				println(err.Error())
				wifiClient.ConnectWifi()
				mqttClient.ConnectBroker()
			default:
				println(err.Error())
			}
		}

	}
}
