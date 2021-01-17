package main

import "fmt"

func main() {
	sliceOfInt := []int{10, 20, 30, 40, 50}

	// развернули slice
	calculate(sliceOfInt...)
}

func calculate(numbers ...int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}
