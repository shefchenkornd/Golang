package main

import "fmt"

// ключевое слово `type` вводит новый тип
// За ним следует имя нового типа (Circle) и ключевое слово struct, которое говорит,
// что мы определяем структуру и список полей внутри фигурных скобок.

type Circle struct {
	x, y, radius float64
}

func main() {
	circle := Circle{1, 3, 5}

	fmt.Println(circle.radius)

	point := &circle

	fmt.Println(point.x)

	fmt.Println(a)
}
