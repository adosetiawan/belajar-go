package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAdo NoKTP = "317xxxxxxxxxxxx"
	fmt.Println(ktpAdo)

	var contoh string = "20000000000"

	var contoh2 = NoKTP(contoh)
	fmt.Println(contoh2)
}
