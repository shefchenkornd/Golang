package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

// Получатель указывается в отдельном списке аргументов между ключевым словом func и именем метода.
// В данном примере метод Abs имеет получателя с типом Vertex и именем v.
func (v Vertex) Abs() float64  {
	return math.Sqrt(v.x + v.y)
}

func main() {
	
	v:= Vertex{3, 7}

	fmt.Println(v.Abs())
}
