package main

import "fmt"

// Using struct for Encapsulation of Behavior
type gopher struct {
	name string
	age  int
}

func main() {
	young := gopher{name: "Phill", age: 25}
	old := gopher{name: "Karl", age: 65}

	young.Jump() // Telling gopher what to do. Logic is encapsulated and hidden from the caller
	old.Jump()

}

func (g gopher) Jump() {
	if g.age < 60 {
		fmt.Println(g.name, "jumping high")
	} else {
		fmt.Println(g.name, "jump a little")
	}
}
