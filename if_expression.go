package main

import "fmt"

func main() {

	name := "Ado se"

	if name == "Ado se" {
		fmt.Println("Hello Ado se")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hello Unknown")
	}

	//sort statement
	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name is acceptable")
	}
}
