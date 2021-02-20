package main

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/dom"
)

type SensorEvent struct {
	TimeStamp   string
	Message     string
	Temperature float32
	Pressure    float32
	Humidity    float32
}

var sensorEvents []SensorEvent

func splitter(this js.Value, args []js.Value) interface{} {
	values := strings.Split(args[0].String(), ",")

	result := make([]interface{}, 0)
	for _, each := range values {
		result = append(result, each)
	}

	return js.ValueOf(result)
}

func handleSensorEvents(channel chan SensorEvent) {
	for {
		event := <-channel
		println("received sensor event")
		sensorEvents = append(sensorEvents, event)

		tableBody := dom.GetElementByID("tbody")

		tr := dom.CreateElement("tr")

		addTd(tr, event.TimeStamp)
		addTd(tr, event.Message)
		addTdf(tr, "%v°C", event.Temperature)
		addTdf(tr, "%v hPa", event.Pressure)
		addTdf(tr, "%v", event.Humidity)

		dom.AppendChild(tableBody, tr)

	}
}

func addTd(tr js.Value, value interface{}) {
	td := dom.CreateElement("td")
	dom.SetInnerHTML(td, value)
	dom.AppendChild(tr, td)
}

func addTdf(tr js.Value, formatString string, value interface{}) {
	td := dom.CreateElement("td")
	dom.SetInnerHTML(td, fmt.Sprintf(formatString, value))
	dom.AppendChild(tr, td)
}

func main() {
	sensorEvents := make(chan SensorEvent)
	go handleSensorEvents(sensorEvents)
	go testAddToTable(sensorEvents)

	wait := make(chan struct{}, 0)
	<-wait
}

func testAddToTable(messages chan SensorEvent) {
	for {
		messages <- SensorEvent{
			TimeStamp:   time.Now().Format(time.RFC3339),
			Message:     "Sensor reading",
			Temperature: 24.32,
			Pressure:    1000.15,
			Humidity:    54.32,
		}
		time.Sleep(3 * time.Second)
	}
}
