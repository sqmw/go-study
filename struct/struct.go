package main

import "fmt"

type person struct {
	name string
	age  int
}

func getPersonByName(name string) *person {
	return &person{name: name}
}

func main() {
	var p *person = new(person)
	fmt.Println(p)
}
