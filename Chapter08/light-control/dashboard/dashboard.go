package dashboard

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/table"
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter08/light-control/login"
)

var doc = tinydom.GetDocument()

type Service struct {
	user          login.UserInfo
	bedroomLights bool
	logoutChannel chan struct{}
}

func New(logout chan struct{}) *Service {
	js.Global().Set("handleMessage", js.FuncOf(handleMessage))
	js.Global().Set("handleOnConnect", js.FuncOf(handleOnConnect))

	return &Service{
		logoutChannel: logout,
	}
}

func (service *Service) ConnectMQTT() {
	println("connecting to mqtt")
	js.Global().
		Get("MQTTconnect").
		Invoke()
}

func handleOnConnect(this js.Value, args []js.Value) interface{} {
	requestStatus()
	return nil
}

func handleMessage(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	println("status message arrived:", message)

	messageParts := strings.Split(message, "#")

	room := messageParts[0]
	component := messageParts[1]

	switch room {
	case "bedroom":
		switch component {
		case "lights":
			doc.
				GetElementById("bedroom-light-status").
				SetInnerHTML(messageParts[2])
		default:
			println("unknown component:", component)
			return nil
		}
	default:
		println("unknown room:", room)
		return nil
	}

	return nil
}

func (service *Service) RenderDashboard(user login.UserInfo) {
	service.user = user

	tinydom.GetWindow().
		PushState(nil, "dashboard", "/dashboard")

	body := doc.GetElementById("body-component")
	div := doc.CreateElement("div").
		SetId("dashboard-component")

	h1 := doc.CreateElement("h1").
		SetInnerHTML("Dashboard")
	h2 := doc.CreateElement("h2").
		SetInnerHTML(fmt.Sprintf("Hello %s", service.user.UserName))

	tableElement := table.New().
		SetHeader(
			"Component",
			"Actions",
			"Status",
		)

	tbody := doc.CreateElement("tbody")

	tr := doc.CreateElement("tr")
	componentNameElement := doc.CreateElement("td").
		SetInnerHTML("Bedroom Lights")
	componentControlElement := doc.CreateElement("td")
	statusElement := doc.CreateElement("td").
		SetId("bedroom-light-status").
		SetInnerHTML("off")

	onButton := input.New(input.ButtonInput).
		SetValue("On").
		AddEventListener("click", js.FuncOf(service.bedroomOn))

	offButton := input.New(input.ButtonInput).
		SetValue("Off").
		AddEventListener("click", js.FuncOf(service.bedroomOff))

	componentControlElement.AppendChildren(onButton, offButton)

	tr.AppendChildren(componentNameElement, componentControlElement, statusElement)

	tbody.AppendChildren(tr)

	tableElement.SetBody(tbody)

	logout := input.New(input.ButtonInput).
		SetValue("logout").
		AddEventListener("click", js.FuncOf(service.logout))

	div.AppendChildren(
		h1,
		h2,
		tableElement.Element,
		tinydom.GetDocument().CreateElement("br"),
		logout,
	)
	body.AppendChild(div)
}

func (service *Service) logout(this js.Value, args []js.Value) interface{} {
	service.logoutChannel <- struct{}{}
	return nil
}

func requestStatus() {
	println("requesting status")
	js.Global().Get("publish").Invoke("home/status-request", "")
}

func (service *Service) bedroomOn(this js.Value, args []js.Value) interface{} {
	if time.Now().After(service.user.LoggedInAt.Add(5 * time.Minute)) {
		println("timeOut: perform logout")
		service.logout(js.ValueOf(nil), nil)
		return nil
	}

	println("turning lights on")
	js.Global().Get("publish").Invoke("home/bedroom/lights", "on")

	service.user.LoggedInAt = time.Now()

	return nil
}

func (service *Service) bedroomOff(this js.Value, args []js.Value) interface{} {
	if time.Now().After(service.user.LoggedInAt.Add(5 * time.Minute)) {
		println("timeOut: perform logout")
		service.logout(js.ValueOf(nil), nil)
		return nil
	}

	println("turning lights off")
	js.Global().Get("publish").Invoke("home/bedroom/lights", "off")

	service.user.LoggedInAt = time.Now()

	return nil
}
