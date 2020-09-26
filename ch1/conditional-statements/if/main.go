package main

func main() {
	if 13%2 == 0 {
		println("13 is even")
	} else {
		println("13 is odd")
	}

	if num := 1; num > 0 {
		println(num, " is positive")
	} else if num > 100 {
		println(num, " has atleast 3 digits")
	} else {
		println(num, "is negative and has an unknown number of digits")
	}
}
