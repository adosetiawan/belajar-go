package main

import "fmt"

func main() {
	// break statement
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
		if i == 5 {
			break
		}
	}
	fmt.Println("-----------------------------------")
	//continue statement
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
