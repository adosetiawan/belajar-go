package main

import (
	"belajar-go/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Dunya")
	fmt.Println(result)

	fmt.Println(helper.Application)         // error
	fmt.Println(helper.Version)             // error
	fmt.Println(helper.SayGoodbye("Dunya")) // error
}
