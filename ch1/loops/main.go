package main

func main() {
	for index := 0; index <= 100; index++ {
		println(index)
	}

	count := 10000
	for {
		if count <= 0 {
			break
		}
		count--
	}
	println("count to 0 successfully")
}
