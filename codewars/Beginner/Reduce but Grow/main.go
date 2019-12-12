package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3,7,1,5}
	sort.Ints(arr)

	result := 1

	for _, v :=range arr {
		result *= v
	}

	fmt.Println(result)
}
