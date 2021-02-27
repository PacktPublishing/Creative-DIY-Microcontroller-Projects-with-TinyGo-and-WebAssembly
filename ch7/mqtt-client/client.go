package mqttclient

import (
	"math/rand"
	"time"

	"github.com/Nerzal/drivers/net/mqtt"
)

type Client interface {
	ConnectBroker() error
	PublishMessage(topic string, message []byte) error
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

func (client *client) PublishMessage(topic string, message []byte) error {
	token := client.mqttClient.Publish(topic, 0, false, message)
	if token.WaitTimeout(time.Second) && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
