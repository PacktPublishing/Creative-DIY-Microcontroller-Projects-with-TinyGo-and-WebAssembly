hello-world-of-things: 
	tinygo flash --target=arduino Chapter01/hello-world-of-things/main.go

writing-to-serial: 
	tinygo flash --target=arduino Chapter03/writing-to-serial/main.go

controlling-keypad: 
	tinygo flash --target=arduino Chapter03/controlling-keypad/main.go

safety-lock-keypad: 
	tinygo flash --target=arduino Chapter03/safety-lock-keypad/main.go

soil-moisture-sensor:
	tinygo flash --target=arduino Chapter04/soil-moisture-sensor/main.go && cat /dev/ttyACM0

water-level-sensor:
	tinygo flash --target=arduino Chapter04/water-level-sensor/main.go && cat /dev/ttyACM0

buzzer-example: 
	tinygo flash --target=arduino Chapter04/buzzer-example/main.go

plant-watering-system:
	tinygo flash --target=arduino Chapter04/plant-watering-system/main.go

soil-moisture-sensor-thresholds:
	tinygo flash --target=arduino Chapter04/soil-moisture-sensor-thresholds/main.go

ultrasonic-distance-sensor-example:
	tinygo flash --target=arduino-nano33 Chapter05/ultrasonic-distance-sensor-example/main.go && putty

hs42561k-example:
	tinygo flash --target=arduino-nano33 Chapter05/hs42561k-example/main.go

hs42561k-spi-example:
	tinygo flash --target=arduino-nano33 Chapter05/hs42561k-spi-example/main.go

hd44780-text-display:
	tinygo flash --target=arduino-nano33 Chapter06/hd44780-text-display/main.go

hd44780-user-input:
	tinygo flash --target=arduino-nano33 Chapter06/hd44780-user-input/main.go

hd44780-cli:
	tinygo flash --target=arduino-nano33 Chapter06/hd44780-cli/main.go

st7735:
	tinygo flash --target=arduino-nano33 Chapter06/st7735/main.go

game:
	tinygo flash --target=arduino-nano33 Chapter06/tinygame/main.go

weather:
	tinygo flash --target=arduino-nano33 Chapter07/weather-station-example/main.go

wasm-app:
	rm -rf Chapter07/html
	mkdir Chapter07/html
	tinygo build -o Chapter07/html/wasm.wasm -target wasm -no-debug Chapter07/weather-app/wasm.go
	cp Chapter07/weather-app/wasm_exec.js Chapter07/html/
	cp Chapter07/weather-app/wasm.js Chapter07/html/
	cp Chapter07/weather-app/index.html Chapter07/html/
	go run Chapter07/wasm-server/main.go

start-mqtt-broker:
	docker start mosquitto

weather-station-mqtt:
	tinygo flash --target=arduino-nano33 Chapter07/weather-station-mqtt/main.go

light-control:
	rm -rf Chapter08/html
	mkdir Chapter08/html
	tinygo build -o Chapter08/html/wasm.wasm -target wasm -no-debug Chapter08/light-control/wasm.go
	cp Chapter08/light-control/wasm_exec.js Chapter08/html/
	cp Chapter08/light-control/wasm.js Chapter08/html/
	cp Chapter08/light-control/dashboard.html Chapter08/html/
	cp Chapter08/light-control/index.html Chapter08/html/
	go run Chapter08/wasm-server/main.go

test:
	tinygo test --tags "arduino_nano33" Chapter05/ultrasonic-distance-sensor/driver_test.go

serial: 
	cat /dev/ttyACM0