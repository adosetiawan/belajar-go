package main

import "fmt"

func random() interface{} {
	return "true"
}
func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	// fmt.Printf("Type: %T\n", resultString)

	// panic: interface conversion: interface {} is string, not int
	// resultInt int := result.(int)
	// fmt.Println(resultInt)
	// fmt.Printf("Type: %T\n", resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("result is string:", value)
	case int:
		fmt.Println("result is int:", value)
	default:
		fmt.Println("unknown type")
	}
}
