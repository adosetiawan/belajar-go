package main

import "fmt"

func main() {
	var name1 = "ado"
	var name2 = "ado"

	fmt.Println(name1 == name2) // true
	fmt.Println(name1 != name2)
	fmt.Println(name1 > name2) // false
	fmt.Println(name1 < name2)

	var nilaiAkhir = 81
	var absensi = 90

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println("lulus:", lulus)
}
