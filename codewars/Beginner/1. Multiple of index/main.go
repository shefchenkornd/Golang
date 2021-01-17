package main

import "fmt"

func main() {
	ints := []int{22, -6, 32, 82, 9, 25}

	fmt.Println(multipleOfIndex(ints))
}

//Return a new array consisting of elements which are multiple of their own index in input array (length > 1).
func multipleOfIndex(ints []int) []int {
	newArray := make([]int, 0)

	for index, value := range ints {
		if index > 0 && value % index == 0 {
			newArray = append(newArray, value)
		}
	}
	return newArray
}
