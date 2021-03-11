package mqttclient

import (
	"math/rand"
	"time"

	"github.com/Nerzal/drivers/net/mqtt"
)

type Client struct {
	mqttBroker   string
	mqttClientID string
	MqttClient   mqtt.Client
}

func New(mqttBroker, clientID string) *Client {
	return &Client{
		mqttBroker:   mqttBroker,
		mqttClientID: clientID,
	}
}

func (client *Client) ConnectBroker() error {
	opts := mqtt.NewClientOptions().
		AddBroker(client.mqttBroker).
		SetClientID(client.mqttClientID + randomString(4))

	client.MqttClient = mqtt.NewClient(opts)

	token := client.MqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (client *Client) PublishMessage(topic string, message []byte, qos uint8, retain bool) error {
	token := client.MqttClient.Publish(topic, qos, retain, message)
	if token.WaitTimeout(time.Second) && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (client *Client) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := client.MqttClient.Subscribe(topic, qos, callback)
	if token.WaitTimeout(time.Second) && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
