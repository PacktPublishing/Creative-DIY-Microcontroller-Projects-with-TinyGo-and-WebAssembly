package main

import (
	"strconv"
	"strings"
	"syscall/js"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/dom"
)

type sensorEvent struct {
	TimeStamp   string
	Message     string
	Temperature float32
	Pressure    float32
	Humidity    float32
}

type alertEvent struct {
	TimeStamp string
	Message   string
	Diff      string
	TimeSpan  string
}

var dataMessages = make(chan sensorEvent)
var alertMessages = make(chan alertEvent)

func main() {
	js.Global().Set("sensorDataHandler", js.FuncOf(sensorDataHandler))
	js.Global().Set("alertHandler", js.FuncOf(alertHandler))

	go handleSensorEvents(dataMessages)
	go handleAlertEvents(alertMessages)
	// go testAddToTable(sensorEvents)

	wait := make(chan struct{}, 0)
	<-wait
}

func handleSensorEvents(channel chan sensorEvent) {
	for {
		event := <-channel
		println("adding sensor event to table")

		tableBody := dom.GetElementByID("tbody-data")
		tr := dom.CreateElement("tr")
		dom.AddTd(tr, event.TimeStamp)
		dom.AddTd(tr, event.Message)
		dom.AddTdf(tr, "%v°C", event.Temperature)
		dom.AddTdf(tr, "%v hPa", event.Pressure)
		dom.AddTdf(tr, "%v", event.Humidity)

		dom.AppendChild(tableBody, tr)
		println("successfully added sensor event to table")
	}
}

func handleAlertEvents(channel chan alertEvent) {
	for {
		event := <-channel
		println("adding sensor event to table")

		tableBody := dom.GetElementByID("tbody-alerts")
		tr := dom.CreateElement("tr")
		dom.AddTd(tr, event.TimeStamp)
		dom.AddTd(tr, event.Message)
		dom.AddTdf(tr, "%s hPa", event.Diff)
		dom.AddTdf(tr, "%s", event.TimeSpan)

		dom.AppendChild(tableBody, tr)
		println("successfully added sensor event to table")
	}
}

func alertHandler(this js.Value, args []js.Value) interface{} {
	println("mqtt: alert message received")

	// payload: message#diff#timespan
	message := args[0].String()

	println("message:", message)
	splittedStrings := strings.Split(message, "#")
	println("splitted string length:", len(splittedStrings))

	alertMessages <- alertEvent{
		TimeStamp: time.Now().Format(time.RFC1123),
		Message:   splittedStrings[0],
		Diff:      splittedStrings[1],
		TimeSpan:  splittedStrings[2],
	}

	return nil
}

func sensorDataHandler(this js.Value, args []js.Value) interface{} {
	println("mqtt: data message received")

	message := args[0].String()

	println("message:", message)
	splittedStrings := strings.Split(message, "#")

	temperature, err := strconv.ParseFloat(splittedStrings[1], 32)
	if err != nil {
		println("failed to parse temperature from message")
		println(message)
	}

	pressure, err := strconv.ParseFloat(splittedStrings[2], 32)
	if err != nil {
		println("failed to parse pressure from message")
		println(message)
	}

	humidity, err := strconv.ParseFloat(splittedStrings[3], 32)
	if err != nil {
		println("failed to parse humidity from message")
		println(message)
	}

	dataMessages <- sensorEvent{
		TimeStamp:   time.Now().Format(time.RFC1123),
		Message:     splittedStrings[0],
		Temperature: float32(temperature),
		Pressure:    float32(pressure),
		Humidity:    float32(humidity),
	}

	return nil
}

func testAddToTable(messages chan sensorEvent) {
	for {
		messages <- sensorEvent{
			TimeStamp:   time.Now().Format(time.RFC3339),
			Message:     "Sensor reading",
			Temperature: 24.32,
			Pressure:    1000.15,
			Humidity:    54.32,
		}
		time.Sleep(3 * time.Second)
	}
}
