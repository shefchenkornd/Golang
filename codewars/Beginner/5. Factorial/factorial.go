package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}

func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
