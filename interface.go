package main

import "fmt"

// Interface definition

// HasName interface requires a GetName method
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Printf("Hello, %s!\n", hasName.GetName())
}

// Implementation of the interface
type Person struct {
	Name string
}
type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Alice"}
	fmt.Println(person)
	SayHello(person)

	animal := Animal{Name: "Capibara"}
	fmt.Println(animal)
	SayHello(animal)
}
