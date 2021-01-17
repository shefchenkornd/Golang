package main

import "fmt"

//                                                                                                  type
//                                                                                                 /
// Interfaces provide a way to specify behavior: "If something can do this, then it can be used here".
//                                                                      |
//                                                                     method
// A convention for naming one-method interfaces in Go is to use the method name an -er suffix.
type jumper interface {
	jump() string  // Method expected to be present in all types that implement this interface
}

type gopher struct {
	name string
	age int
}

func (g gopher) jump() string {
	return "gopher jump"
}

type horse struct {
	name string
	weight int
}

func (h horse) jump() string {
	return "horse jump"
}

func main() {
	jumperList := getList() // Compiler does NOT care about naming, so these could be named somethin else too....
	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}
}

//  jumper -> can be used as return type (The * symbol is not necessary when working with interfaces)
func getList() []jumper {
	phil := &gopher{
		name: "phil",
		age:  25,
	}

	barbaro := &horse{
		name:   "barbaro",
		weight: 2000,
	}

	return []jumper{phil, barbaro} // All types part of this slice MUST respond to jump()
}