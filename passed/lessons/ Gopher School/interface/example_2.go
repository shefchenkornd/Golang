package main

import "fmt"

type greeter interface{
	greet(string) string
}

type russian struct {}
type american struct {}

func (r *russian) greet(name string) string  {
	return fmt.Sprintf("Привет, %v", name)
}

func (a *american) greet(name string) string  {
	return fmt.Sprintf("Hello, %v", name)
}

func greet(g greeter, name string)  {
	fmt.Println(g.greet(name))
}

func main() {
	greet(&russian{}, "Петя")
	greet(&american{}, "Bill")

}
