package keypad

import (
	"machine"
)

// Keypad is a driver for 4x4 keypads
type Keypad struct {
	inputEnabled bool
	lastColumn   int
	lastRow      int
	columns      [4]machine.Pin
	rows         [4]machine.Pin
	mapping      [4][4]string
}

func (keypad *Keypad) Configure(c4, c3, c2, c1, r4, r3, r2, r1 machine.Pin) {
	inputConfig := machine.PinConfig{Mode: machine.PinInputPullup}
	c4.Configure(inputConfig)
	c3.Configure(inputConfig)
	c2.Configure(inputConfig)
	c1.Configure(inputConfig)

	keypad.columns = [4]machine.Pin{c4, c3, c2, c1}

	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	r4.Configure(outputConfig)
	r3.Configure(outputConfig)
	r2.Configure(outputConfig)
	r1.Configure(outputConfig)

	r4.High()
	r3.High()
	r2.High()
	r1.High()

	keypad.rows = [4]machine.Pin{r4, r3, r2, r1}

	keypad.mapping = [4][4]string{
		{"1", "2", "3", "A"},
		{"4", "5", "6", "B"},
		{"7", "8", "9", "C"},
		{"*", "0", "#", "D"},
	}

	keypad.inputEnabled = true
	keypad.lastColumn = -1
	keypad.lastRow = -1
}

func (keypad *Keypad) GetButton() string {
	row, column := keypad.GetIndices()
	if row == -1 && column == -1 {
		return ""
	}

	return keypad.mapping[row][column]
}

func (keypad *Keypad) GetIndices() (int, int) {
	for rowIndex := range keypad.rows {
		rowPin := keypad.rows[rowIndex]
		rowPin.Low()

		for columnIndex := range keypad.columns {
			columnPin := keypad.columns[columnIndex]

			if !columnPin.Get() && keypad.inputEnabled {
				keypad.inputEnabled = false

				keypad.lastColumn = columnIndex
				keypad.lastRow = rowIndex

				return keypad.lastRow, keypad.lastColumn
			}

			if columnPin.Get() &&
				columnIndex == keypad.lastColumn &&
				rowIndex == keypad.lastRow &&
				!keypad.inputEnabled {
				keypad.inputEnabled = true
			}
		}

		rowPin.High()
	}

	return -1, -1
}
