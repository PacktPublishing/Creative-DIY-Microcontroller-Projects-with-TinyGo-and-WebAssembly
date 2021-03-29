package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"
)

var buttonPressed bool

const enemySize = 8
const bulletSize = 4
const width = 128
const height = 160

var highscore int = 0
var currentScore int = 0

var (
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
	black = color.RGBA{0, 0, 0, 255}
)

func main() {
	buttonPin := machine.D9
	buttonPin.Configure(machine.PinConfig{Mode: machine.PinInput})

	buttonPowerPin := machine.D10
	buttonPowerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	buttonPin.High()

	updateHighscore(0)

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
	})

	resetPin := machine.D6
	dcPin := machine.D5
	csPin := machine.D7
	backLightPin := machine.D2

	display := st7735.New(machine.SPI0, resetPin, dcPin, csPin, backLightPin)
	display.Configure(st7735.Config{})

	go checkButton(buttonPin)

	for {
		display.FillScreen(black)
		updateGame(display)
	}
}

func checkButton(buttonPin machine.Pin) {
	for {
		if buttonPin.Get() {
			buttonPressed = true
		}

		time.Sleep(20 * time.Millisecond)
	}
}

func updateHighscore(score int) {
	if score <= highscore && score != 0 {
		return
	}

	highscore = score

	println(fmt.Sprintf("  TinyInvader  Highscore: %d", highscore))
}

func updateGame(display st7735.Device) {
	var enemyPosX, enemyPosY int16
	enemyPosY = height - enemySize

	var bulletPosY int16

	shotFired := false
	canFire := true
	currentScore = 0

	for {
		if buttonPressed {
			buttonPressed = false

			if canFire {
				shotFired = true
				canFire = false
			}
		}

		if shotFired {
			bulletPosY = updateBullet(display, bulletPosY)

			if bulletPosY > height {
				shotFired = false
				canFire = true
				bulletPosY = 0
			}

			if enemyPosX >= 54 && enemyPosX <= 64 {
				if enemyPosY >= bulletPosY && enemyPosY <= bulletPosY+bulletSize {
					currentScore++

					display.FillRectangle(enemyPosX-1, enemyPosY, enemySize, enemySize, black)

					enemyPosY = height - enemySize
					enemyPosX = 0

					updateHighscore(currentScore)
				}
			}
		}

		enemyPosX, enemyPosY = updateEnemy(display, enemyPosX, enemyPosY)
		if enemyPosY < enemySize {
			return
		}

		display.FillRectangle(0, 4, width, 1, green)
		display.FillRectangle(58, 0, 6, 6, green)

		time.Sleep(12 * time.Millisecond)
	}
}

func updateBullet(display st7735.Device, posY int16) int16 {
	display.FillRectangle(58, posY-2, bulletSize, 2, black)
	display.FillRectangle(58, posY, bulletSize, bulletSize, green)
	return posY + 2
}

func updateEnemy(display st7735.Device, posX, posY int16) (int16, int16) {
	var clearX, clearY, clearWidth int16

	clearX = posX - 1
	clearY = posY
	clearWidth = 1
	if posX == 0 {
		clearY = posY + enemySize
		clearX = width - enemySize
		clearWidth = enemySize
	}

	display.FillRectangle(clearX, clearY, clearWidth, enemySize, black)
	display.FillRectangle(posX, posY, enemySize, enemySize, red)

	posX++
	if posX > width-enemySize {
		posX = 0
		posY -= enemySize
	}

	return posX, posY

}
