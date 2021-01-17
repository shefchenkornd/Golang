package main

import "fmt"

/**
написать функцию `Concat`, которая получает несколько слайсов и
склеивает их в один длинный
{ {1,2,3}, {4,5} } => {  }
 */
func Concat(slices [][]int) []int {
	var result []int

	for _, slice := range slices {
		result = append(result, slice ...)
	}

	return result
}

func main() {
	slices := [][]int{{1, 2}, {3, 4}}
	fmt.Println(Concat(slices))
}
