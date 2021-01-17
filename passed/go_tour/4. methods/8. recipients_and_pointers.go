package main

import (
	"fmt"
)

type Vertex struct {
	x, y float64
}

// Вы можете объявлять методы, где в качестве получателей выступают указатели.
// Это означает, что получатель объявлен как *T для некоторого типа T.
func (v *Vertex) Scale(f float64)  {
	v.x *= f
	v.y *= f
}

func main() {
	v := Vertex{2,8}
	v.Scale(10)

	fmt.Println(v)
}
