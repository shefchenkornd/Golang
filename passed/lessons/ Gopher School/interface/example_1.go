package main

import "fmt"

type animal interface {
	makeSound()
}

type Cat struct{}
type Dog struct{}

func (c *Cat) makeSound() {
	fmt.Println("meow")
}

func (d *Dog) makeSound() {
	fmt.Println("gaw")
}

func main() {
	// поскольку Cat{} и Dog{} реализуют интерфейс animal, то мы можем указать тип интерфейса animal у переменных
	var c, d animal = &Cat{}, &Dog{}

	fmt.Println(c)

	c.makeSound()
	d.makeSound()
}
