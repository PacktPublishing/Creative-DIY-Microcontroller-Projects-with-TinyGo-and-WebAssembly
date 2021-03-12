package main

import (
	"time"

	"github.com/Nerzal/tinydom"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter08/light-control/dashboard"
	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter08/light-control/login"
)

var window = tinydom.GetWindow()
var loginService *login.Service
var dashboardService dashboard.Service

type userInfo struct {
	loggedIn   bool
	userName   string
	loggedInAt time.Time
}

var loginState userInfo

func main() {
	loginState = userInfo{}

	loginChannel := make(chan (string), 1)

	loginService = login.NewService(loginChannel)
	loginService.RenderLogin()
	go onLogin(loginChannel)

	dashboardService = dashboard.New()

	wait := make(chan struct{}, 0)
	<-wait
}

func onLogin(channel chan (string)) {
	for {
		userName := <-channel
		println(userName, "logged in!")

		loginState.userName = userName
		loginState.loggedIn = true
		loginState.loggedInAt = time.Now()

		removeLoginComponent()
		dashboardService.ConnectMQTT()
		dashboardService.RenderDashboard()
	}
}

func removeLoginComponent() {
	doc := tinydom.GetDocument()
	doc.GetElementById("body-component").
		RemoveChild(doc.GetElementById("login-component"))
}

func logout() {
	loginState = userInfo{}
	window.PushState(nil, "login", "")

}
