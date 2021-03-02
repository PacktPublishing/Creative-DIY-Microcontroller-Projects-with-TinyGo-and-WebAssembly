package main

import (
	"time"
)

func main() {
	print("starting ")
	print("program\n")

	for {
		println("Hello World")
		time.Sleep(1 * time.Second)
	}
}
