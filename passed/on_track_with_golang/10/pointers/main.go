package main

import "fmt"

type gopher struct {
	name string
	age int
	isAdult bool
}

// indicates a pointer to the original gopher.
func validateAge(g *gopher)  {
	g.isAdult = g.age >= 18
}

// Using the & operator, a reference to the existing memory address is assigned to the new varable.
func main() {
	// The & operator returns a Pointer to this new struct.
	var phil = &gopher{
		name:    "Phil",
		age:     35,
		isAdult: false,
	}
	validateAge(phil) // здесь передаю указатель на адрес переменной в памяти
	fmt.Println(phil)
}
