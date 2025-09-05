package main

import "fmt"

func main() {
	// Declare and initialize an array of integers
	var months = [12]string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
	var slice = months[3:6] // Slicing the array to get a subset
	fmt.Println(slice)

	var slice1 = months[10:] // Slicing from index 10 to the end
	fmt.Println(slice1)

	var slice2 = months[:4] // Slicing from the start to index 4 (exclusive)
	fmt.Println(slice2)

	var triwulan1 = months[0:3]  // Slicing from index 0 to 3 (exclusive)
	var triwulan2 = months[3:6]  // Slicing from index 3 to 6 (exclusive)
	var triwulan3 = months[6:9]  // Slicing from index 6 to 9 (exclusive)
	var triwulan4 = months[9:12] // Slicing from index 9 to 12 (exclusive)
	var smester1 = months[:6]    // Slicing from index 0 to 6 (exclusive)
	var smester2 = months[6:]    // Slicing from index 6 to the end

	fmt.Println("Triwulan 1:", triwulan1)
	fmt.Println("Triwulan 2:", triwulan2)
	fmt.Println("Triwulan 3:", triwulan3)
	fmt.Println("Triwulan 4:", triwulan4)

	fmt.Println("Smester 1:", smester1)
	fmt.Println("Smester 2:", smester2)

	// Modifying the first element of the smester2 slice
	smester2[0] = "July Baru"
	fmt.Println(smester2)
	fmt.Println("---------------------------------------------")

	//append slice
	appendBulan := append(smester2, "Bulan Baru")
	fmt.Println(smester2)
	appendBulan[1] = "Agustus Baru"
	fmt.Println(appendBulan)

	days := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	fmt.Println("---------------------------------------------")
	//make slice parameter length and capacity
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ado"
	newSlice[1] = "Rahmad"

	fmt.Println(newSlice)
	//megnhitung panjang
	fmt.Println(len(newSlice))
	//menghitung kapasitas
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice3 := append(newSlice2, "Eko")
	fmt.Println(newSlice3)
	fmt.Println(len(newSlice3))
	fmt.Println(cap(newSlice3))

	newSlice4 := append(newSlice3, "Eko")
	fmt.Println(newSlice4)
	fmt.Println(len(newSlice4))
	fmt.Println(cap(newSlice4))
	//copy slice

	newSlice5 := make([]string, len(newSlice4), cap(newSlice4))
	copy(newSlice5, newSlice4)
	fmt.Println(newSlice5)
}
