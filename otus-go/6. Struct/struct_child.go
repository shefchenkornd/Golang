package main

import "fmt"

type Base struct {}
func (b Base) Name() string {
 return "Base"
}
func (b Base) Say() {
	fmt.Println(b.Name())
}

type Child struct {
	Base
	name string
}
func (c Child) Name() string {
	return "Child"
}

func main() {
	c := Child{}
	c.Say() // Base
}
