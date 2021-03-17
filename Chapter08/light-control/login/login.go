package login

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/form"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/label"
)

const user = "tinygo"
const password = "secure1234"

var doc = tinydom.GetDocument()

type Service struct {
	channel chan string
}

func NewService(channel chan string) *Service {
	return &Service{channel: channel}
}

func (service *Service) RenderLogin() {
	tinydom.GetWindow().PushState(nil, "login", "/login")

	div := doc.CreateElement("div").
		SetId("login-component")
		
	h1 := doc.CreateElement("h1").
		SetInnerHTML("Login")

	loginForm := form.New()

	userNameLabel := label.
		New().
		SetFor("userName").
		SetInnerHTML("UserName:")

	userName := input.
		New(input.TextInput).
		SetId("userName")

	passwordLabel := label.
		New().
		SetFor("password").
		SetInnerHTML("Password:")

	password := input.
		New(input.PasswordInput).
		SetId("password")

	login := input.New(input.ButtonInput).
		SetValue("login").
		AddEventListener("click", js.FuncOf(service.onClick)).
		AddEventListener("keypress", js.FuncOf(service.onKeyPress))

	loginForm.AppendChildrenBr(
		userNameLabel,
		userName,
		passwordLabel,
		password,
		login,
	)

	div.AppendChildren(h1, loginForm.Element)
	
	body := doc.GetElementById("body-component")
	body.AppendChildren(div)
}

func (service *Service) onClick(this js.Value, args []js.Value) interface{} {
	service.login()

	return nil
}

func (service *Service) onKeyPress(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		println("keyPress event with no event should not happen")
		return nil
	}

	event := tinydom.Event{Value: args[0]}
	if event.Key() == "Enter" {
		service.login()
	}

	return nil
}

func (service *Service) login() {
	userElem := input.FromElement(doc.GetElementById("userName"))
	userName := userElem.Value()

	if userName != user {
		tinydom.GetWindow().Alert("Invalid username or password")
		return
	}

	passwordElem := input.FromElement(doc.GetElementById("password"))
	passwordInput := passwordElem.Value()

	if passwordInput != password {
		tinydom.GetWindow().Alert("Invalid username or password")
		return
	}

	go func() {
		service.channel <- userName
	}()
}
