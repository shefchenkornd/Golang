package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1      = Vertex{1, 2}
	v2      = Vertex{x: 10}
	v3      = Vertex{x: 0, y: 30}
	pointer = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3,pointer)
}
