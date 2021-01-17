package main

import (
	"fmt"
)

func Pic2(dx, dy  int) [][]uint8  {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = make([]uint8, dx)
	}

	for y := range result {
		for x := range result[y] {
			result[y][x] = uint8(x*y)
		}
	}

	return  result
}

func main() {
	multidimensionalSlice := Pic2(12, 25)
	for y := range multidimensionalSlice {
		fmt.Printf("Тип %T - значение: %v\n", multidimensionalSlice[y], multidimensionalSlice[y])
	}
}
