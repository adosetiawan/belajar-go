package main

import (
	"fmt"
)

func main() {
	//call multiple value with ignored value
	_, lastName := getFullName()
	fmt.Println(lastName)
	//call function with multiple return value
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	//call function without parameter and return value
	generateGenap()
	//call function with parameter
	printArray([5]int{10, 20, 30, 40, 50 * 10})
	//call function return value
	fmt.Println("function return value:", getNama(2))
	//named return value
	firstName, middleName, _ := getCompleteName()
	fmt.Println("named return value:", firstName, middleName)

	//call variadic function
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println("variadic function:", total)

	//call variadic function with slice
	sumSliceData := []int{10, 20, 30, 40, 50}
	totalSlice := sumAllSlice(sumSliceData)
	totalSliceAll := sumAll(sumSliceData...)
	fmt.Println("variadic function with slice all:", totalSliceAll)
	fmt.Println("variadic function with slice:", totalSlice)

	// call function value
	Goodbye := getGoodbye
	fmt.Println("function value:", Goodbye("Ado"))
	//call function as parameter
	sayHelloWithFilter("Ado", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)

	//call anonymous function
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	//call factorial function
	result := factorialLoop(10)
	fmt.Println("factorial loop:", result)
	//call factorial recursive function
	resultRecursive := factorialRecursive(10)
	fmt.Println("factorial recursive:", resultRecursive)
}
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

// factorial function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// anonymous function
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

type Blacklist func(string) bool

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}
func spamFilter(name string) string {
	if name == "anjing" {
		return "****"
	}
	return name
}

// function value
func getGoodbye(name string) string {
	return "Goodbye " + name
}

// variadic function with slice
func sumAllSlice(numbers []int) int {
	total := 0
	// weith for range
	for _, number := range numbers {
		total += number
	}

	return total
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	// weith for range
	for _, number := range numbers {
		total += number
	}

	return total
}

// named return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ado"
	middleName = "Rahmad"
	lastName = "Sundara"
	return firstName, middleName, lastName
}

// return multiple value
func getFullName() (string, string) {
	return "Ado", "Rahmad"
}

// function without parameter and return value
func generateGenap() {
	var genap [5]int
	var j = 2
	for i := 0; i < 5; i++ {
		genap[i] = j
		j += 2
	}
	fmt.Println(genap)
}

// function with parameter
func printArray(arrays [5]int) {
	for i := 0; i < len(arrays); i++ {
		fmt.Println(arrays[i])
	}
}

// function return value
func getNama(number int) string {
	var nama string
	if number == 1 {
		nama = "Eko"
	} else {
		nama = "Eko patrion"
	}
	return nama
}
