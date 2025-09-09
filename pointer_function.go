package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Tasikmalaya", "Jawa Barat", "sunda empire"}
	// ChangeAddressToIndonesia(address1)
	ChangeAddressToIndonesia(&address1)
	fmt.Println(address1) // address1 tidak berubah karena address di function ChangeAddressToIndonesia adalah pass by value
}
