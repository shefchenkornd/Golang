package main

import "fmt"

/**
... - pack operator, он собирает все параметры типа Type
в слайс, и он же их распаковывает.
ТОЛЬКО ПОСЛЕДНИЙ ПАРАМЕТР функции может быть вариадическим.
*/
func Foo(format string, a ...int) int {
	fmt.Println(format)

	s := 0
	for _, i := range a {
		s += i
	}

	return s
}

func main() {
	s := Foo("int", 1, 2, 3, 4, 5)
	fmt.Printf("%d \n", s)

	sl := []int{8, 9}
	a := []int{10, 11, 12}
	sl = append(sl, a...) // разворачиваем слайс на элементы
	fmt.Println(sl)
}
