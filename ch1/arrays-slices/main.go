package main

func main() {
	names := [4]string{"Alice", "Bob", "Charlie", "Denise"}

	for index, name := range names {
		println("index:", index, "name:", name)
	}

	namesSlice := []string{"Nerzal"}

	for _, name := range names {
		namesSlice = append(namesSlice, name)
	}

	for index, name := range namesSlice {
		println("index:", index, "name:", name)
	}
}
