package main

import (
	"machine"
	"time"
)

// Left
const pos0DutyCycle = 1500 * time.Microsecond
const pos0RemainingPeriod = 18500 * time.Microsecond

// Right
const pos2DutyCycle = 2000 * time.Microsecond
const pos2RemainingPeriod = 18000 * time.Microsecond

var inputEnabled = true
var lastColumn = -1
var lastRow = -1
var columns []machine.Pin
var rows []machine.Pin
var mapping [4][4]string
var servoPosition = 0

func main() {
	initialize()

	machine.InitPWM()
	servoPin := machine.PWM{Pin: machine.D11}
	servoPin.Configure()

	const password = "133742"
	enteredPassword := ""

	for {
		rowIndex, columnIndex := getIndices()
		if rowIndex != -1 && columnIndex != -1 {
			println("Button: ", mapping[columnIndex][rowIndex])

			enteredPassword += mapping[columnIndex][rowIndex]
		}

		if len(enteredPassword) == len(password) {
			if password == enteredPassword {
				println("Success")
				enteredPassword = ""
				close(servoPin)
			} else {
				println("Fail")
				println("Entered Password: ", enteredPassword)
				enteredPassword = ""
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func open(servoPin machine.PWM) {
	for servoPosition = 0; servoPosition >= 1; servoPosition-- {
		servoPin.Pin.High()
		time.Sleep(pos2DutyCycle)
		servoPin.Pin.Low()
		time.Sleep(pos2RemainingPeriod)

		time.Sleep(100 * time.Millisecond)
	}
}

func close(servoPin machine.PWM) {
	for servoPosition = 0; servoPosition < 180; servoPosition++ {
		servoPin.Pin.High()
		time.Sleep(pos0DutyCycle)
		servoPin.Pin.Low()
		time.Sleep(pos0RemainingPeriod)
	}
}

func initialize() {
	inputConfig := machine.PinConfig{Mode: machine.PinInputPullup}
	c4 := machine.D2
	c4.Configure(inputConfig)
	c3 := machine.D3
	c3.Configure(inputConfig)
	c2 := machine.D4
	c2.Configure(inputConfig)
	c1 := machine.D5
	c1.Configure(inputConfig)

	columns = []machine.Pin{c4, c3, c2, c1}

	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	r4 := machine.D6
	r4.Configure(outputConfig)
	r3 := machine.D7
	r3.Configure(outputConfig)
	r2 := machine.D8
	r2.Configure(outputConfig)
	r1 := machine.D9
	r1.Configure(outputConfig)

	r4.High()
	r3.High()
	r2.High()
	r1.High()

	rows = []machine.Pin{r4, r3, r2, r1}

	mapping = [4][4]string{
		{"1", "2", "3", "A"},
		{"4", "5", "6", "B"},
		{"7", "8", "9", "C"},
		{"*", "0", "#", "D"},
	}
}

func getIndices() (int, int) {
	for rowIndex := range rows {
		rowPin := rows[rowIndex]

		rowPin.Low()

		for columnIndex := range columns {
			columnPin := columns[columnIndex]

			if !columnPin.Get() && inputEnabled {
				inputEnabled = false

				lastColumn = columnIndex
				lastRow = rowIndex

				return lastRow, lastColumn
			}

			if columnPin.Get() &&
				columnIndex == lastColumn &&
				rowIndex == lastRow &&
				!inputEnabled {
				inputEnabled = true
			}
		}

		rowPin.High()
	}

	return -1, -1
}