package main

import "fmt"

func main() {
	defer three()
	defer two()
	one()
}

func one()  {
	fmt.Println("one")
}

func two()  {
	fmt.Println("two")
}

func three()  {
	fmt.Println("three")
}
