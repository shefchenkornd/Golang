package main

import "fmt"

// Задача:
// Чи́сла Фибона́ччи — элементы числовой последовательности
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89 ...
// в которой первые два числа равны либо 1 и 1, либо 0 и 1, а каждое последующее число равно сумме двух предыдущих чисел.

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	slice := []int{0, 1}

	return func() int {
		sum := 0
		length := len(slice) - 2

		for _, value := range slice[length:] {
			sum += value
		}
		slice = append(slice, sum)

		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}