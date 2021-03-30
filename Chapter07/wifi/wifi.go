package wifi

import (
	"machine"
	"time"

	"github.com/Nerzal/drivers/wifinina"
)

type Client interface {
	Configure() error
	CheckHardware() error
	ConnectWifi()
}

type client struct {
	ssid     string
	password string
	spi      machine.SPI
	wifi     *wifinina.Device
}

func New(ssid, password string) Client {
	return &client{
		spi:      machine.NINA_SPI,
		ssid:     ssid,
		password: password,
	}
}

func (client *client) Configure() error {
	err := client.spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})
	if err != nil {
		return err
	}

	// Using "old" wifinina driver use this
	// client.wifi = &wifinina.Device{
	// 	SPI:   client.spi,
	// 	CS:    machine.NINA_CS,
	// 	ACK:   machine.NINA_ACK,
	// 	GPIO0: machine.NINA_GPIO0,
	// 	RESET: machine.NINA_RESETN,
	// }

	wifiDevice := wifinina.NewSPI(
		client.spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN,
	)

	client.wifi = wifiDevice

	client.wifi.Configure()

	// Needs some time to configure, before connection to wifi can be established
	time.Sleep(5 * time.Second)

	return nil
}

func (client *client) CheckHardware() error {
	firmwareVersion, err := client.wifi.GetFwVersion()
	if err != nil {
		return err
	}

	println("firmware version: ", firmwareVersion)

	result, err := client.wifi.ScanNetworks()
	if err != nil {
		return err
	}

	for i := 0; i < int(result); i++ {
		ssid := client.wifi.GetNetworkSSID(i)
		println("ssid: ", ssid, "id:", i)
	}

	return nil
}

func (client *client) ConnectWifi() {
	println("trying to connect to network: ", client.ssid)

	client.connect()

	for {
		time.Sleep(1 * time.Second)

		status, err := client.wifi.GetConnectionStatus()
		if err != nil {
			println("error:", err.Error())
		}

		println("status:", status.String())

		if status == wifinina.StatusConnected {
			break
		}

		if status == wifinina.StatusConnectFailed || status == wifinina.StatusDisconnected {
			client.connect()
		}
	}

	ip, _, _, err := client.wifi.GetIP()
	if err != nil {
		println("could not get ip address:", err.Error())
	}

	println("connected to wifi. got ip:", ip.String())
}

func (client *client) connect() error {
	if client.password == "" {
		err := client.wifi.SetNetwork(client.ssid)
		return err
	}

	return client.wifi.SetPassphrase(client.ssid, client.password)
}
