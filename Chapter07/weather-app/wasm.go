package main

import (
	"strings"
	"syscall/js"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter07/dom"
)

type sensorEvent struct {
	TimeStamp   string
	Message     string
	Temperature string
	Pressure    string
	Humidity    string
}

type alertEvent struct {
	TimeStamp string
	Message   string
	Diff      string
	TimeSpan  string
}

var sensorEvents = make(chan sensorEvent)
var alertEvents = make(chan alertEvent)

func main() {
	js.Global().Set("sensorDataHandler", js.FuncOf(sensorDataHandler))
	js.Global().Set("alertHandler", js.FuncOf(alertHandler))

	go handleSensorEvents()
	go handleAlertEvents()
	// go testAddToTable(sensorEvents)

	wait := make(chan struct{}, 0)
	<-wait
}

func alertHandler(this js.Value, args []js.Value) interface{} {
	println("mqtt: alert message received")

	// payload: message#diff#timespan
	message := args[0].String()

	println("message:", message)
	splittedStrings := strings.Split(message, "#")
	println("splitted string length:", len(splittedStrings))

	alertEvents <- alertEvent{
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

	sensorEvents <- sensorEvent{
		TimeStamp:   time.Now().Format(time.RFC1123),
		Message:     splittedStrings[0],
		Temperature: splittedStrings[1],
		Pressure:    splittedStrings[2],
		Humidity:    splittedStrings[3],
	}

	return nil
}

func handleSensorEvents() {
	for {
		event := <-sensorEvents
		println("adding sensor event to table")

		tableBody := dom.GetElementByID("tbody-data")
		tr := dom.CreateElement("tr")
		dom.AddTd(tr, event.TimeStamp)
		dom.AddTd(tr, event.Message)
		dom.AddTdf(tr, "%s°C", event.Temperature)
		dom.AddTdf(tr, "%s hPa", event.Pressure)
		dom.AddTdf(tr, "%s", event.Humidity)

		dom.AppendChild(tableBody, tr)
		println("successfully added sensor event to table")
	}
}

func handleAlertEvents() {
	for {
		event := <-alertEvents
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

func testAddToTable(messages chan sensorEvent) {
	for {
		messages <- sensorEvent{
			TimeStamp:   time.Now().Format(time.RFC3339),
			Message:     "Sensor reading",
			Temperature: "24.32",
			Pressure:    "1000.15",
			Humidity:    "54.32",
		}
		time.Sleep(3 * time.Second)
	}
}
