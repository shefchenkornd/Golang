package main

import "fmt"

// Функции как тип (functions as a type golang)
//func main() {
//	t := calculate(100)
//
//	// в переменной t хранится анонимная функция, пример, ниже
//	//t := func(k int) int {
//	//	return k * 100
//	//}
//
//
//	z:=t(5)
//	fmt.Println(z)
//}
//
//func calculate(i int) func(int)int {
//	fmt.Println(i)
//
//	return func(k int) int {
//		return k * 100
//	}
//}

func main() {
	sliceOfNUmbers := []int{10, 20, 30, 40, 50}

	result := calculate2(sliceOfNUmbers, func(n int) int {
		return n * 3
	})

	fmt.Println("result =", result)
}

func calculate2(numbers []int, justFunc func(int) int) int {
	result := 0

	for _, v := range numbers {
		fmt.Println("for in v =", v)

		result += justFunc(v)
	}

	return result
}
