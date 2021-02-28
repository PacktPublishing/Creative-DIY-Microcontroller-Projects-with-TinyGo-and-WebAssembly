package mqttclient

import (
	"math/rand"
	"time"

	"github.com/Nerzal/drivers/net/mqtt"
)

type Client interface {
	ConnectBroker() error
	PublishMessage(topic string, message []byte, qos uint8, retain bool) error
}

type client struct {
	mqttBroker   string
	mqttClientID string
	mqttClient   mqtt.Client
}

func New(mqttBroker string) Client {
	return &client{
		mqttBroker: mqttBroker,
	}
}

func (client *client) ConnectBroker() error {
	opts := mqtt.NewClientOptions().
		AddBroker(client.mqttBroker).
		SetClientID(client.mqttClientID + randomString(4))

	client.mqttClient = mqtt.NewClient(opts)

	token := client.mqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (client *client) PublishMessage(topic string, message []byte, qos uint8, retain bool) error {
	token := client.mqttClient.Publish(topic, qos, retain, message)
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
