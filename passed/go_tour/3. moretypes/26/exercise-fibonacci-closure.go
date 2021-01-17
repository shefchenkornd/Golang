package main

import "fmt"

// Задание:
// Давайте повеселимся с функциями.
// Реализуйте функцию fibonacci, которая возвращает функцию (замыкание), которая возвращает последовательные числа Фибоначчи (0, 1, 1, 2, 3, 5, ...).

// fibonacci is a function that returns
// a function that returns an int.
func fibonaci() func() int {
	sl := []int{0}
	return func() int {
		result := 0
		length := len(sl)
		if length == 1 {
			sl = append(sl, 1)
			return 0
		}

		index := length - 2
		for _, v := range sl[index:] {
			result += v
		}
		sl = append(sl, result)

		return sl[length-1]
	}
}

func main() {
	f := fibonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
