package main

import "fmt"

type Address struct {
	City, Province, Country string
}
type Man struct {
	Name string
}

func (man *Man) Maried() {
	man.Name = "Mr." + man.Name
}
func main() {
	var man1 Man = Man{"ado"}
	man1.Maried()

	fmt.Println(man1.Name)
}
