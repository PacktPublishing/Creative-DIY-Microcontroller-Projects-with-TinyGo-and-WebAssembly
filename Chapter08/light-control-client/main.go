package main

import (
	"fmt"
	"machine"
	"strings"
	"time"

	"github.com/Nerzal/drivers/net/mqtt"
	mqttclient "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter07/mqtt-client"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter07/wifi"
)

const ssid = ""
const password = ""

const bedroomLight = machine.D4

var bedroomLightStatus = false

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

	err = mqttClient.Subscribe("home/control", 0, HandleActionMessage)
	if err != nil {
		printError("could not subsribe to topic", err)
	}

	err = mqttClient.Subscribe("home/status-request", 0, HandleStatusRequestMessage)
	if err != nil {
		printError("could not subsribe to topic", err)
	}

	// use for test purposes
	// go func() {
	// 	for {
	// 		mqttClient.PublishMessage("home/control", []byte("bedroom#lights#on"), 0, false)
	// 		time.Sleep(500 * time.Millisecond)
	// 		println("published message")
	// 	}
	// }()

	println("subscribed to topic, waiting for messages")

	select {}
}

func HandleStatusRequestMessage(client mqtt.Client, message mqtt.Message) {
	println("handling status request")
	reportStatus(client)
	message.Ack()
}

func reportStatus(client mqtt.Client) {
	status := "off"
	if bedroomLightStatus {
		status = "on"
	}

	token := client.Publish("home/status", 0, false, fmt.Sprintf("bedroom#lights#%s", status))
	if token.Wait() && token.Error() != nil {
		println(token.Error())
	}
}

// room # module # action
func HandleActionMessage(client mqtt.Client, message mqtt.Message) {
	println("handling incoming message")
	payload := string(message.Payload())
	splittedString := strings.Split(payload, "#")

	if len(splittedString) != 3 {
		println("invalid message:", payload)
		message.Ack()
		return
	}

	println(
		"room:",
		splittedString[0],
		"module:",
		splittedString[1],
		"action:",
		splittedString[2],
	)

	switch splittedString[0] {
	case "bedroom":
		controlBedroom(
			client,
			splittedString[1],
			splittedString[2],
		)
	default:
		println("invalid room:", payload)
	}

	message.Ack()
}

func controlBedroom(client mqtt.Client, module, action string) {
	switch module {
	case "lights":
		switch action {
		case "on":
			controlBedroomlights(client, true)
		case "off":
			controlBedroomlights(client, false)
		default:
			println("unknown action:", action)

		}
	default:
		println("invalid module:", module)
	}
}

func controlBedroomlights(client mqtt.Client, action bool) {
	if action {
		bedroomLight.High()
		bedroomLightStatus = true
	} else {
		bedroomLight.Low()
		bedroomLightStatus = false
	}

	reportStatus(client)
}

func printError(message string, err error) {
	for {
		println(message, err.Error())
		time.Sleep(time.Second)
	}
}
