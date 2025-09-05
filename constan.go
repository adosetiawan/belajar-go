package main

import "fmt"

func main() {
	const firstName = "Ado"
	const lastName = "Setiawan"

	//trying change constant value
	// firstName = "Tidak bisa di ubah"
	// lastName = "Tidak bisa di ubah"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//,ultiple constant
	const (
		alamat1 = "Jl. Mawar"
		alamat2 = "Jl. Melati"
		alamat3 = "Jl. Kenanga"
	)
	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println(alamat3)
}
