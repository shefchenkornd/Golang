package main

import "fmt"

func main() {
	// array - это структура с ФИКСИРОВАННОЙ длиной
	var array [7]int

	array[0] = 0
	array[3] = 3
	array[6] = 6

	fmt.Println(array)

	for i, v := range  array {
		fmt.Println("index", i, "value ", v)
	}

	// объявили и проинициализировали массив той длиной,
	// которая равна количеству элементов в нём
	dynamicArray := [...]string{"a", "b", "c", "d", "f", "g", "j"}

	// длина массива
	fmt.Println(len(dynamicArray))
	fmt.Println(dynamicArray)
}
