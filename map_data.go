package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Tom"
	// person["age"] = "20"
	// fmt.Println(person)

	//make map parameter 1: key, parameter 2: value
	person := map[string]string{
		"name":    "ado",
		"address": "Indonesia",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//access key empty
	fmt.Println(person["title"])

	//delete map by key
	delete(person, "address")
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "Learning Go"
	book["author"] = "John Doe"
	book["price"] = "10000"
	fmt.Println(book)
}
