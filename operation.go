package main

import "fmt"

// main demonstrates basic arithmetic operations in Go.
// It performs addition, multiplication, and subtraction operations
// on integer variables and prints the results to console.
//
// Variables:
// a - first operand with value 10
// b - second operand with value 10
// g - third operand with value 2
// h - fourth operand with value 2
// c - result of addition (a + b)
// d - result of multiplication (a * b)
// f - result of subtraction (c - d)
//
// The function prints individual results and a compound arithmetic expression.
func main() {
	var a = 10
	var b = 10
	var g = 2
	var h = 2

	var c = a + b
	var d = a * b
	var f = c - d
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

	fmt.Println(a + b - h*g)

	//augmanted assigment
	a += 2 // a = a + 2
	fmt.Println(a)
	a += 4 // a = a + 2
	fmt.Println(a)

	b *= 2 // b = b * 2
	fmt.Println(b)

	//unary operation
	a++ // a = a + 1
	fmt.Println(a)
	a-- // a = a - 1
	fmt.Println(a)

	//operasi perbandingan

	var benar bool = true
	var salah bool = false

	fmt.Println(benar == salah)
	fmt.Println(benar != salah)
	fmt.Println(benar && salah)
	fmt.Println(benar || salah)

}
