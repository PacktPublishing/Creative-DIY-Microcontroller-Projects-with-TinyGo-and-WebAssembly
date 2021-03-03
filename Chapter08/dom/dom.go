package dom

import (
	"fmt"
	"syscall/js"

	"github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/Chapter07/dom"
)

func GetDocument() js.Value {
	return js.Global().Get("document")
}

func CreateElement(tag string) js.Value {
	document := GetDocument()
	return document.Call("createElement", tag)
}

func GetElementByID(id string) js.Value {
	document := GetDocument()
	return document.Call("getElementById", id)
}

func AppendChild(parent js.Value, child js.Value) {
	parent.Call("appendChild", child)
}

func SetInnerHTML(object js.Value, value interface{}) {
	object.Set("innerHTML", value)
}

func AddTd(tr js.Value, value interface{}) {
	td := CreateElement("td")
	SetInnerHTML(td, value)
	AppendChild(tr, td)
}

func AddTdf(tr js.Value, formatString string, value interface{}) {
	td := CreateElement("td")
	SetInnerHTML(td, fmt.Sprintf(formatString, value))
	AppendChild(tr, td)
}

func GetInputValue(id string) string {
	passwordInput := dom.GetElementByID(id)
	return passwordInput.Get("value").String()
}

func Alert(message string) {
	js.Global().
		Get("window").
		Call("alert", message)
}

func ExportFunction(jsFunction string, goFunction func(this js.Value, args []js.Value) interface{}) {
	js.Global().Set(jsFunction, js.FuncOf(goFunction))
}

// https://developer.mozilla.org/en/docs/Web/API/History_API
func PushState(state interface{}, title, URL string) {
	js.Global().
		Get("window").
		Get("history").
		Call("pushState", state, title, URL)
}

func ReplaceState(state interface{}, title, URL string) {
	js.Global().
		Get("window").
		Get("history").
		Call("replaceState", state, title, URL)
}
