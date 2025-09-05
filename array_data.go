package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ado"
	names[1] = "Setiawan"
	names[2] = "Raziq"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[2] = "Ganteng" // akan error karena index array hanya sampai 2
	fmt.Println(names[2])

	// deklarasi dan inisialisasi array
	var values = [3]int{
		90, 80, 70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	//function array

	var valuesDynamic = [...]int{
		90, 80, 70, 100, 110,
	}
	fmt.Println(valuesDynamic, len(valuesDynamic))
}
