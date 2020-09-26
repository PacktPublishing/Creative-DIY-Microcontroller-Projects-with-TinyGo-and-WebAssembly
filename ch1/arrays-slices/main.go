package main

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Denise"}

	for index, name := range names {
		println("index:", index, "name:", name)
	}

	for _, name := range names {
		println("name:", name)
	}

	var invited []string
	for _, name := range names {
		invited = append(invited, name)
	}

	println("len:", len(invited), "capacity:", cap(invited))

	var notInvited [2]string
	j := 0
	for i := range names {
		if i%2 == 0 {
			continue
		}

		name := names[i]
		notInvited[j] = name
		j++
	}

	for _, name := range notInvited {
		println("Not invited:", name)
	}
}
