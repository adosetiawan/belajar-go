package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	//overflow jika nilai melebihi batas tipe data
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "ado"
	var a uint8 = name[0]
	var eString = string(a)
	fmt.Println(a)
	fmt.Println(eString)
}
