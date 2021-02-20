# Programming-Microcontrollers-and-WebAssembly-with-TinyGo

Programming Microcontrollers and WebAssembly with TinyGo, published by Packt

## Required Software

[Git](https://git-scm.com/)

[Go](https://golang.org/)

### Windows

#### AVR Toolchain

[Toolchain](https://www.microchip.com/mplab/avr-support/avr-and-arm-toolchains-c-compilers)
[AVR Dude](http://download.savannah.gnu.org/releases/avrdude/)
[GNU Make](http://gnuwin32.sourceforge.net/packages/make.htm)

## Recommended Software

[VSCode](https://code.visualstudio.com/)

[VSCode Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.go)

[VSCode TinyGo Extension](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo)

## Required Hardware

[Arduino UNO](https://store.arduino.cc/arduino-uno-rev3)

[Arduino Nano 33 IoT](https://store.arduino.cc/arduino-nano-33-iot)

## Troubleshooting

### Windows

#### exit status 3221225781

When tinygo flash returns an error like this: 

> error: failed to flash C:\Users\Enrico\AppData\Local\Temp\tinygo393394635\main.hex: exit status 3221225781
 
This is most likely due to a missing .dll file.
Install this to get the dll you need: https://sourceforge.net/projects/libusb-win32/

## DataSheets

[MAX7219](https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf)

[HD44780](https://cdn.shopify.com/s/files/1/1509/1638/files/HD44780_1602_Blaues_LCD_Display_mit_Serielle_Schnittstelle_I2C_Bundle_Datenblatt_AZ-Delivery_Vertriebs_GmbH.pdf?v=1591601507)

[ST7735](https://cdn.shopify.com/s/files/1/1509/1638/files/1_8_inch_OLED_Datenblatt_04323b18-84e6-4e7b-bf7d-3fa56a308f66.pdf?633464727103137069)

[bmp280](https://www.bosch-sensortec.com/media/boschsensortec/downloads/datasheets/bst-bmp280-ds001.pdf)