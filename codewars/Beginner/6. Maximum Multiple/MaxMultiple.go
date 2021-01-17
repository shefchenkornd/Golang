package main

import "fmt"

func main() {
	divisor := 2
	bound := 9

	fmt.Println(MaxMultiple(divisor, bound))
}

func MaxMultiple(d, b int) int {
	var result int = 1

	result = b % d

	if result == 0 {
		return b
	}

	return MaxMultiple(d, b-1)
}
