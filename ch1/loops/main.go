package main

func main() {
	for i := 0; i <= 100; i++ {
		println(i)
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
