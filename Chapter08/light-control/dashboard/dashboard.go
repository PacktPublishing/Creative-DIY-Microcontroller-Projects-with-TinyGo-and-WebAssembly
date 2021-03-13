package dashboard

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/table"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter08/light-control/login"
)

var doc = tinydom.GetDocument()

type Service struct {
	user          login.UserInfo
	lastAction    time.Time
	logoutChannel chan bool
}

func New(logout chan bool) Service {
	return Service{
		logoutChannel: logout,
	}
}

func (service *Service) ConnectMQTT() {
	println("connecting to mqtt")
	js.Global().Get("MQTTconnect").Invoke()
}

func (service *Service) RenderDashboard(user login.UserInfo) {
	service.user = user

	tinydom.GetWindow().
		PushState(nil, "dashboard", "/dashboard")

	body := doc.GetElementById("body-component")
	div := doc.CreateElement("div").
		SetId("dashboard-component")

	h1 := doc.CreateElement("h1").SetInnerHTML("Dashboard")
	h2 := doc.CreateElement("h2").SetInnerHTML(fmt.Sprintf("Hello %s", service.user.UserName))

	tableElement := table.New().
		SetHeader("Component", "Actions")

	tbody := doc.CreateElement("tbody")

	tr := doc.CreateElement("tr")
	componentNameElement := doc.CreateElement("td").
		SetInnerHTML("Bedroom Lights")
	componentControlElement := doc.CreateElement("td")

	onButton := input.New(input.ButtonInput).
		SetValue("On").
		AddEventListener("click", js.FuncOf(service.bedroomOn))

	offButton := input.New(input.ButtonInput).
		SetValue("Off").
		AddEventListener("click", js.FuncOf(service.bedroomOff))

	componentControlElement.AppendChildren(onButton, offButton)

	tr.AppendChildren(componentNameElement, componentControlElement)

	tbody.AppendChildren(tr)

	tableElement.SetBody(tbody)

	logout := input.New(input.ButtonInput).
		SetValue("logout").
		AddEventListener("click", js.FuncOf(service.logout))

	div.AppendChildren(h1, h2, tableElement.Element, tinydom.GetDocument().CreateElement("br"), logout)
	body.AppendChild(div)
}

func (service *Service) logout(this js.Value, args []js.Value) interface{} {
	service.logoutChannel <- true
	return nil
}

func (service *Service) bedroomOn(this js.Value, args []js.Value) interface{} {
	println("turning lights on")

	// room # module # action
	js.Global().Get("publish").Invoke("home-control", "bedroom#lights#on")
	return nil
}

func (service *Service) bedroomOff(this js.Value, args []js.Value) interface{} {
	println("turning lights off")
	js.Global().Get("publish").Invoke("home-control", "bedroom#lights#off")
	return nil
}
