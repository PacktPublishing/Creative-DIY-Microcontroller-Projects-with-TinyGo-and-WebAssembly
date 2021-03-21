package main

func main() {
	resultChannel := make(chan bool)
	errChannel := make(chan error)

	select {
	case result := <-resultChannel:
		println(result)
	case err := <-errChannel:
		println(err.Error())
	}
}
