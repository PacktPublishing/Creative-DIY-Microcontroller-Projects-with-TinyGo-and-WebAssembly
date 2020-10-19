package main

func main() {
	i := 1337

	pointer := &i
	println(*pointer)

	*pointer = 23
	println(i)

	person := Person{Age: 18}
	addYear(person)
	println("Age is:", person.Age)

	addYearPointer(&person)
	println("Age is:", person.Age)
}

type Person struct {
	Age int
}

func addYear(p Person) {
	p.Age++
}

func addYearPointer(p *Person) {
	p.Age++
}
