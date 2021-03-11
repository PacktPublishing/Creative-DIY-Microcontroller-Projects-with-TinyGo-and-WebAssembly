package main

import (
	"machine"
	"strings"
	"time"

	"github.com/Nerzal/drivers/net/mqtt"
	mqttclient "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter07/mqtt-client"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter07/wifi"
)

const ssid = ""
const password = ""

const bedroomLight = machine.D10

func main() {
	time.Sleep(5 * time.Second)

	bedroomLight.Configure(machine.PinConfig{Mode: machine.PinOutput})

	wifiClient := wifi.New(ssid, password)

	println("configuring nina wifi chip")
	err := wifiClient.Configure()
	if err != nil {
		printError("could not configure wifi client", err)
	}

	println("checking firmware")
	wifiClient.CheckHardware()

	wifiClient.ConnectWifi()

	mqttClient := mqttclient.New("tcp://192.168.2.102:1883", "lightControl")
	println("connecting to mqtt broker")
	err = mqttClient.ConnectBroker()
	if err != nil {
		printError("could not configure mqtt", err)
	}
	println("connected to mqtt broker")

	mqttClient.Subscribe("lights", 1, HandleActionMessage)

	for {
		time.Sleep(time.Minute)
	}

}

// room # module # action
func HandleActionMessage(client mqtt.Client, message mqtt.Message) {
	payload := string(message.Payload())
	splittedString := strings.Split(payload, "#")

	if len(splittedString) != 3 {
		println("invalid message:", payload)
		message.Ack()
		return
	}

	println("room:", splittedString[0], "module:", splittedString[1], "action:", splittedString[2])

	switch splittedString[1] {
	case "bedroom":
		controlBedroom(splittedString[1], splittedString[2])
	default:
		println("invalid room:", payload)
	}

	message.Ack()
}

func controlBedroom(module, action string) {
	switch module {
	case "lights":
		switch action {
		case "on":
			controlBedroomlights(true)
		case "off":
			controlBedroomlights(false)
		default:
			println("unknown action:", action)

		}
	default:
		println("invalid module:", module)
	}
}

func controlBedroomlights(action bool) {
	if action {
		bedroomLight.High()
	} else {
		bedroomLight.Low()
	}
}

func printError(message string, err error) {
	for {
		println(message, err.Error())
		time.Sleep(time.Second)
	}
}
