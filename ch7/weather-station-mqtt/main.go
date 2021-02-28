package main

import (
	"fmt"
	"machine"
	"time"

	mqttclient "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/mqtt-client"
	weatherstation "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/weather-station"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/wifi"
	"tinygo.org/x/drivers/bme280"
	"tinygo.org/x/drivers/wifinina"
)

// insert your network ssid
const ssid = "changeMe"

// insert your network password, leave empty, if you access an open network
const password = "changeME"

var (
	temperature float64
	pressure    float64
	humidity    float64
)

func printError(message string, err error) {
	println(message, err.Error())
	time.Sleep(time.Second)
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

	wifiClient.CheckHardware()

	wifiClient.ConnectWifi()

	mqttClient := mqttclient.New("tcp://192.168.2.102:1883")
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

func publishSensorData(mqttClient mqttclient.Client, wifiClient wifi.Client, weatherStation weatherstation.Service) {
	for {
		time.Sleep(time.Minute)
		println("publishing sensor data")

		tempString, pressString, humidityString := weatherStation.GetFormattedReadings(temperature, pressure, humidity)
		message := []byte(fmt.Sprintf("%s#%s#%s#%s", "sensor readings", tempString, pressString, humidityString))

		err := mqttClient.PublishMessage("weather/data", message, 0, true)
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

func publishAlert(mqttClient mqttclient.Client, wifiClient wifi.Client, weatherStation weatherstation.Service) {
	// pressure loss of more than 6 hPa / 3h -> possible storm incoming
	// 							  2 hpa / 1h -> possible storm incoming
	// source http://www.bohlken.net/airpressure2.htm
	for {
		time.Sleep(time.Hour)

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
