package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Status        bool
}

func main() {
	var ado Customer

	fmt.Println(ado)

	ado.Name = "Ado"
	ado.Address = "Indonesia"
	ado.Age = 20
	ado.Status = false
	fmt.Println(ado)
	fmt.Println(ado.Name)
	fmt.Println(ado.Address)
	fmt.Println(ado.Age)
	fmt.Println(ado.Status)

	rahmad := Customer{
		Name:    "Rahmad",
		Address: "Indonesia",
		Age:     21,
	}
	fmt.Println(rahmad)

	budi := Customer{"Budi", "Indonesia", 22, true}
	fmt.Println(budi)

	//struct method
	rully := Customer{
		Name:    "Rully",
		Address: "Indonesia",
		Age:     22,
		Status:  true,
	}
	rully.sayHello()
	budi.sayHello()
	rahmad.sayHello()
}

//method (function with receiver)
func (customer Customer) sayHello() {
	fmt.Println("Hello", "My Name is", customer.Name)
}
