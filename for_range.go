package main

import "fmt"

func main() {
	names := []string{"Andi", "Budi", "Tono"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Index", i, "=", names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println("Without Index", "=", name)
	}

}
