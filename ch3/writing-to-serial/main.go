package main

import "time"

func main() {
	println("starting program")

	for {
		println("Hello World")
		time.Sleep(1 * time.Second)
	}
}
