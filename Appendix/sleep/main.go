package main

import "time"

func main() {
	println("started program")

	for {
		println("work on this goroutine")
		time.Sleep(50 * time.Millisecond)
	}
}
