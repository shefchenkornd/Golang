package main

import "fmt"

type Cats struct {
	name string
	age int
	happiness float64
}

func main()  {
	bob := Cats{
		name:      "Bob",
		age:       3,
		happiness: 0.56,
	}

	fmt.Println(bob)
	fmt.Println("Bob age is", bob.age)
	fmt.Println("Bob function is", bob.test())
}

func (cat *Cats) test() float64  {
	return cat.happiness * float64(cat.age)
}