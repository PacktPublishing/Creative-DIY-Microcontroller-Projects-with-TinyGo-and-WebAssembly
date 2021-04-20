# Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly

Creative DIY Microcontroller Projects with TinyGo and WebAssembly, published by Packt

This is the code repository for [Creative DIY Microcontroller Projects with TinyGo and WebAssembly](https://www.packtpub.com/in/iot-hardware/creative-diy-microcontroller-projects-with-tinygo-and-webassembly?utm_source=github&utm_medium=repository&utm_campaign=), published by Packt.

<a href="https://www.packtpub.com/in/iot-hardware/creative-diy-microcontroller-projects-with-tinygo-and-webassembly?utm_source=github&utm_medium=repository&utm_campaign="><img src="https://static.packt-cdn.com/products/9781800560208/cover/smaller" alt="" height="256px" align="right"></a>

**A practical guide to building embedded applications for low-powered devices, IoT, and home automation**

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

## Required Go packages

```shell
go get -u tinygo.org/x/drivers
go get -u github.com/eclipse/paho.mqtt.golang
```

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

### MacOS

If you have problems to find out, under which name the Arduino registers on your computer, try the following command: 

> ls /dev | grep usb

You are going to see an output like:

> /dev/tty.usbmodem132408

This is the device you can put in the `--p` flag when using `tinygo flash` command

## DataSheets

[MAX7219](https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf)

[HD44780](https://cdn.shopify.com/s/files/1/1509/1638/files/HD44780_1602_Blaues_LCD_Display_mit_Serielle_Schnittstelle_I2C_Bundle_Datenblatt_AZ-Delivery_Vertriebs_GmbH.pdf?v=1591601507)

[ST7735](https://cdn.shopify.com/s/files/1/1509/1638/files/1_8_inch_OLED_Datenblatt_04323b18-84e6-4e7b-bf7d-3fa56a308f66.pdf?633464727103137069)

[bme280](http://www.embeddedadventures.com/datasheets/BME280.pdf)

## MQTT Broker

If you start the broker for the first time run the following command.

docker run -it --name mosquitto \
--net=host \
--restart=always \
-p 1883:1883 \
-p 9001:9001 \
-v ~/go/src/github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter07/mosquitto/config/mosquitto.conf:/mosquitto/config/mosquitto.conf:ro \
eclipse-mosquitto

 For later starts just use:

 > docker start mosquitto

## What is this book about?
While often considered a fast and compact programming language, Go usually creates large executables that are difficult to run on low-memory or low-powered devices such as microcontrollers or IoT. TinyGo is a new compiler that allows developers to compile their programs for such low-powered devices. As TinyGo supports all the standard features of the Go programming language, you won't have to tweak the code to fit on the microcontroller.
This book is a hands-on guide packed full of interesting DIY projects that will show you how to build embedded applications. You will learn how to program sensors and work with microcontrollers such as Arduino UNO and Arduino Nano IoT 33. The chapters that follow will show you how to develop multiple real-world embedded projects using a variety of popular devices such as LEDs, 7-segment displays, and timers. Next, you will progress to build interactive prototypes such as a traffic lights system, touchless hand wash timer, and more. As you advance, you'll create an IoT prototype of a weather alert system and display those alerts on the TinyGo WASM dashboard. Finally, you will build a home automation project that displays stats on the TinyGo WASM dashboard.
By the end of this microcontroller book, you will be equipped with the skills you need to build real-world embedded projects using the power of TinyGo.

This book covers the following exciting features:
Discover a variety of TinyGo features and capabilities while programming your embedded devices. Explore how to use display devices to present your data. Focus on how to make TinyGo interact with multiple sensors for sensing temperature, humidity, and pressure. Program hardware devices such as Arduino Uno and Arduino Nano IoT 33 using TinyGo. Understand how TinyGo works with GPIO, ADC, I2C, SPI, and MQTT network protocols. Build your first TinyGo IoT and home automation prototypes. Integrate TinyGo in modern browsers using WebAssembly

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1800560206) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
func main() {
    blocker := make(chan bool, 1)
    <-blocker
    println("this gets never printed")
}
```

**Following is what you need for this book:**
If you are a Go developer who wants to program low-powered devices and hardware such as Arduino UNO and Arduino Nano IoT 33, or if you are a Go developer who wants to extend your knowledge of using Go with WebAssembly while programming Go in the browser, then this book is for you. Go hobbyist programmers who are interested in learning more about TinyGo by working through the DIY projects covered in the book will also find this hands-on guide useful.

With the following software and hardware list you can run all code files present in the book (Chapter 1-8).
### Software and Hardware List
| No. | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1 | Go 1.15.x or newer | Windows, Mac OS X, and Linux (Any) |
| 2 | Visual Studio Code | Windows, Mac OS X, and Linux (Any) |

## Code in Action
Please visit the following link to check the CiA videos:  https://bit.ly/3cYZOh4

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](https://static.packt-cdn.com/downloads/9781800560208_ColorImages.pdf)

### Related products
* Mastering Go - Second Edition [[Packt]](https://www.packtpub.com/product/mastering-go-second-edition/9781838559335) [[Amazon]](https://www.amazon.com/dp/1838559337)

* Hands-On Software Engineering with Golang [[Packt]](https://www.packtpub.com/product/hands-on-software-engineering-with-golang/9781838554491) [[Amazon]](https://www.amazon.com/dp/1838554491)


## Get to Know the Author
**Tobias Theel** works as Technical Lead and DevOps for German fintech startup fino and since 2020 also for regtech startup ClariLab. As a software architect and particular expert for Go and TinyGo alongside C# and Java, he is iSAQB certified. Theel prepared the base for the SPI support on the Arduino UNO, wrote the tinydom library, which can be used to manipulate the DOM while being fully compatible with TinyGo. He also contributes to the TinyGo drivers repository by adding drivers for new devices. Theel is among the top 10% answerers in C# & Unity3D, as well as the top 20% answerers in .NET, Go, and Visual Studio on Stack Overflow. As a speaker at tech talks and participant in hackathons, Theel shares his knowledge of software development.

### Suggestions and Feedback
[Click here](https://docs.google.com/forms/d/e/1FAIpQLSdy7dATC6QmEL81FIUuymZ0Wy9vH1jHkvpY57OiMeKGqib_Ow/viewform) if you have any feedback or suggestions.
