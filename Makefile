hello-world-of-things: 
	tinygo flash --target=arduino ch1/hello-world-of-things/main.go

writing-to-serial: 
	tinygo flash --target=arduino ch3/writing-to-serial/main.go

controlling-keypad: 
	tinygo flash --target=arduino ch3/controlling-keypad/main.go

safety-lock-keypad: 
	tinygo flash --target=arduino ch3/safety-lock-keypad/main.go

soil-moisture-sensor:
	tinygo flash --target=arduino ch4/soil-moisture-sensor/main.go && cat /dev/ttyACM0

water-level-sensor:
	tinygo flash --target=arduino ch4/water-level-sensor/main.go && cat /dev/ttyACM0

buzzer-example: 
	tinygo flash --target=arduino ch4/buzzer-example/main.go

plant-watering-system:
	tinygo flash --target=arduino ch4/plant-watering-system/main.go

soil-moisture-sensor-thresholds:
	tinygo flash --target=arduino ch4/soil-moisture-sensor-thresholds/main.go

ultrasonic-distance-sensor-example:
	tinygo flash --target=arduino-nano33 ch5/ultrasonic-distance-sensor-example/main.go && putty

hs42561k-example:
	tinygo flash --target=arduino-nano33 ch5/hs42561k-example/main.go

hs42561k-spi-example:
	tinygo flash --target=arduino-nano33 ch5/hs42561k-spi-example/main.go

hd44780-text-display:
	tinygo flash --target=arduino-nano33 ch6/hd44780-text-display/main.go

hd44780-user-input:
	tinygo flash --target=arduino-nano33 ch6/hd44780-user-input/main.go

test:
	tinygo test --tags "arduino_nano33" ch5/ultrasonic-distance-sensor/driver_test.go