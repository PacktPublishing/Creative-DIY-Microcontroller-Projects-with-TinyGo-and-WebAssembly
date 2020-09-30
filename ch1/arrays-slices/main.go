package main

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Denise"}

	for index, name := range names {
		println("index:", index, "name:", name)
	}

	names = append(names, "Nerzal")

	for index, name := range names {
		println("index:", index, "name:", name)
	}
}
