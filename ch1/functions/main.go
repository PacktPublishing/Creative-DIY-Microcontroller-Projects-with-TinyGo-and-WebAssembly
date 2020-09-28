package main

import "errors"

func main() {
	result, err := FizzBuzz(4)
	if err != nil {
		println(err.Error())
	}

	println(result)
}

func FizzBuzz(number int) (string, error) {
	if number < 0 {
		return "", errors.New("negative numbers are not allowed")
	}

	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz", nil
	}

	if number%3 == 0 {
		return "Fizz", nil
	}

	if number%5 == 0 {
		return "Buzz", nil
	}

	return "", nil
}
