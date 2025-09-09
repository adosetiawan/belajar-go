package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Interface definition
func main() {
	// address1 := Address{"Tasikmalaya", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)
	// fmt.Println("--------------------")

	var address1 Address = Address{"Tasikmalaya", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"

	fmt.Println(address1) // address1 berubah karena address2 adalah pointer dari address1
	fmt.Println(address2) // address2 adalah pointer dari address1

}
