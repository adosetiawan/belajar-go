package main

import "fmt"

func main() {
	// for loop with index and value
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("Selesai")
}
