package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Interface definition
func main() {
	var address1 Address = Address{"Tasikmalaya", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"

	fmt.Println(address1) // address1 berubah karena address2 adalah pointer dari address1
	fmt.Println(address2) // address2 adalah pointer dari address1
	fmt.Println("--------------------")
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // mengubah pointer address2
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
