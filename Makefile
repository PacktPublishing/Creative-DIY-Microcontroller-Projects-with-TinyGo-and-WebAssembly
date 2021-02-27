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

hd44780-cli:
	tinygo flash --target=arduino-nano33 ch6/hd44780-cli/main.go

st7735:
	tinygo flash --target=arduino-nano33 ch6/st7735/main.go

game:
	tinygo flash --target=arduino-nano33 ch6/tinygame/main.go

weather:
	tinygo flash --target=arduino-nano33 ch7/weather-station-example/main.go

clean-wasm:
	rm -rf ch7/html
	mkdir ch7/html

wasm_exec:
	cp ch7/weather-app/wasm_exec.js ch7/html/

wasm-app: clean-wasm wasm_exec
	tinygo build -o ch7/html/wasm.wasm -target wasm -no-debug ch7/weather-app/wasm.go
	cp ch7/weather-app/wasm.js ch7/html/
	cp ch7/weather-app/index.html ch7/html/
	go run ch7/wasm-server/main.go

start-mqtt-broker:
	docker start mosquitto

weather-station-mqtt:
	tinygo flash --target=arduino-nano33 ch7/weather-station-mqtt/main.go

test:
	tinygo test --tags "arduino_nano33" ch5/ultrasonic-distance-sensor/driver_test.go

serial: 
	cat /dev/ttyACM0