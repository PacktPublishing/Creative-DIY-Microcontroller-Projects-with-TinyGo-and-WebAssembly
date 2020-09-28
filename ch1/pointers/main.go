package main

func main() {
	i := 1337

	pointer := &i
	println(*pointer)

	*pointer = 23
	println(i)
}
