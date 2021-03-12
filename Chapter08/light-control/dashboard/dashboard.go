package dashboard

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/table"
)

var doc = tinydom.GetDocument()

type Service struct{}

func New() Service {
	return Service{}
}

func (service *Service) ConnectMQTT() {
	println("connecting to mqtt")
	js.Global().Get("MQTTconnect").Invoke()
}

func (service *Service) RenderDashboard() {
	tinydom.GetWindow().PushState(nil, "dashboard", "/dashboard")

	body := doc.GetElementById("body-component")
	div := doc.CreateElement("div").
		SetId("dashboard-component")

	h1 := doc.CreateElement("h1").SetInnerHTML("Dashboard")

	tableElement := table.New().
		SetHeader("Component", "Actions")

	tbody := doc.CreateElement("tbody")

	tr := doc.CreateElement("tr")
	componentNameElement := doc.CreateElement("td").SetInnerHTML("Bedroom Lights")
	componentControlElement := doc.CreateElement("td")

	onButton := input.New(input.ButtonInput).
		SetValue("On").
		AddEventListener("click", js.FuncOf(bedroomOn))
	offButton := input.New(input.ButtonInput).
		SetValue("Off").
		AddEventListener("click", js.FuncOf(bedroomOff))

	componentControlElement.AppendChildren(onButton, offButton)

	tr.AppendChildren(componentNameElement, componentControlElement)

	tbody.AppendChildren(tr)

	tableElement.SetBody(tbody)

	div.AppendChildren(h1, tableElement.Element)
	body.AppendChild(div)
}

func bedroomOn(this js.Value, args []js.Value) interface{} {
	println("turning lights on")

	// room # module # action
	js.Global().Get("publish").Invoke("home-control", "bedroom#lights#on")
	return nil
}

func bedroomOff(this js.Value, args []js.Value) interface{} {
	println("turning lights off")
	js.Global().Get("publish").Invoke("home-control", "bedroom#lights#off")
	return nil
}
