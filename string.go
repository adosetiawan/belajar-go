package main

import "fmt"

func main() {
	fmt.Println("Ado", len("Ado"))
	fmt.Println("Ado Setiawan", len("Ado Setiawan"))
	fmt.Println("Ado"+" middle "+"Setiawan", len("Ado"+" middle "+"Setiawan"))

	fmt.Println("Ado Setiawan"[0])
	fmt.Println("Ado Setiawan"[1])
	fmt.Println("Ado Setiawan"[2])
}
