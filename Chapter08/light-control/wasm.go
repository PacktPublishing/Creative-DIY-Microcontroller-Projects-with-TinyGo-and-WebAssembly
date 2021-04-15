package main

import (
	"time"

	"github.com/Nerzal/tinydom"
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter08/light-control/dashboard"
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter08/light-control/login"
)

var window = tinydom.GetWindow()
var loginService *login.Service
var loginState login.UserInfo
var dashboardService *dashboard.Service

func main() {
	loginState = login.UserInfo{}

	loginChannel := make(chan string, 1)

	loginService = login.NewService(loginChannel)
	loginService.RenderLogin()
	go onLogin(loginChannel)

	logoutChannel := make(chan struct{}, 1)
	go onLogout(logoutChannel)

	dashboardService = dashboard.New(logoutChannel)

	wait := make(chan struct{}, 0)
	<-wait
}

func onLogin(channel chan string) {
	for {
		userName := <-channel
		println(userName, "logged in!")

		loginState.UserName = userName
		loginState.LoggedIn = true
		loginState.LoggedInAt = time.Now()

		removeLoginComponent()
		dashboardService.ConnectMQTT()
		dashboardService.RenderDashboard(loginState)
	}
}

func removeLoginComponent() {
	doc := tinydom.GetDocument()
	doc.GetElementById("body-component").
		RemoveChild(doc.GetElementById("login-component"))
}

func removeDashboardComponent() {
	doc := tinydom.GetDocument()
	doc.GetElementById("body-component").
		RemoveChild(doc.GetElementById("dashboard-component"))
}

func onLogout(channel chan struct{}) {
	for {
		<-channel
		println("handling logout event")
		removeDashboardComponent()
		loginState = login.UserInfo{}

		loginService.RenderLogin()
	}
}
