package main

func main() {
	blocker := make(chan bool, 1)
	<-blocker
	println("this gets never printed")
}
