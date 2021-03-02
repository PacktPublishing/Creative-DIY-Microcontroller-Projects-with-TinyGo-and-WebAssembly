package dom

import (
	"fmt"
	"syscall/js"
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
