package main

import (
	"syscall/js"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter08/dom"
)

func main() {
	js.Global().Set("login", js.FuncOf(login))

	dom.PushState(nil, "login", "/login")

	wait := make(chan struct{}, 0)
	<-wait
}

// export router
func router(event string) {
	println("router called")
}

func login(this js.Value, args []js.Value) interface{} {
	user := dom.GetInputValue("user")
	password := dom.GetInputValue("password")

	if user != "NoobyGames" || password != "TinyGo" {
		dom.Alert("Invalid username or password")
		return nil
	}

	println("Successfully logged in")
	dom.PushState(nil, "dashboard", "/dashboard")

	return nil
}
