package main

func main() {
	value := 1337

	switch {
	case value == 10:
		println("value is 10")
	case value < 0:
		println("value is negative")
	case value > 0:
		println("value is positive")
	}

	name := "Nerzal"
	switch name {
	case "Nerzal":
		println("Name was Nerzal")
	case "Olaf":
		println("Name was Olaf")
	}

	boolValue := true

	switch boolValue {
	case false:
		println("value was false")
	default:
		println("value was true")
	}
}
