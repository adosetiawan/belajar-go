package main

import (
	"fmt"
)

func main() {

	name := "suzana mahardiyansyah"

	switch name {

	case "danis":
		fmt.Println("Hello danis")
	case "budi":
		fmt.Println("Hello budi")
	default:
		fmt.Println("Hi, boleh kenalan")
	}
	fmt.Println("-------------------------------------")
	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length > 5:
		fmt.Println("Name is long")
	default:
		fmt.Println("Name is acceptable")
	}
}
