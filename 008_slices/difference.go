package main

import (
	"fmt"
)

/**
разница в том что простые типы данных, типа int, string, bool, array копируются, а
слайсы передаются по ссылке
вот так!!!
*/
func main() {
	var myArray = [5]int{10, 20, 30}
	var mySlice = []int{100, 200, 300}

	fmt.Printf("Array before DoSomethingArray %v\n", myArray)
	fmt.Printf("Slice before DoSomethingSlice %v\n", mySlice)
	fmt.Println("\n")

	DoSomethingArray(myArray)
	DoSomethingSlice(mySlice)
	fmt.Println("\n")

	fmt.Printf("Array after DoSomethingArray %v\n", myArray)
	fmt.Printf("Slice after DoSomethingSlice %v\n", mySlice)
}

func DoSomethingArray(array [5]int) {
	for i, _ := range array {
		array[i] += 500
	}

	fmt.Printf("Array after finish func %v\n", array)
}

func DoSomethingSlice(slice []int) {
	for i, _ := range slice {
		slice[i] += 1000
	}

	fmt.Printf("Slice after finish func %v\n", slice)
}