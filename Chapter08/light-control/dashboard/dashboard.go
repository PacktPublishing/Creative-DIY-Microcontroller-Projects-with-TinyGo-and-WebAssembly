package dashboard

import (
	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/table"
)

var doc = tinydom.GetDocument()

type Service struct{}

func New() Service {
	return Service{}
}

func (service *Service) RenderDashboard() {
	tinydom.GetWindow().PushState(nil, "dashboard", "/dashboard")

	body := doc.GetElementById("body-component")
	div := doc.CreateElement("div").
		SetId("dashboard-component")

	h1 := doc.CreateElement("h1").SetInnerHTML("Dashboard")

	tableElement := table.New().
		SetHeader("Component", "Actions").
		SetBody("component-control")

	tbody := tableElement.FindChildNode("tbody")
	tr := doc.CreateElement("tr")

	tbody.AppendChildren(tr)

	div.AppendChildren(h1, tableElement.Element)
	body.AppendChild(div)
}
