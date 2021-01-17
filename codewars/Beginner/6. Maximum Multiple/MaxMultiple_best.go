package main

import "fmt"

func main() {
	divisor := 21
	bound := 92

	fmt.Println(MaxMultiple2(divisor, bound))
}

func MaxMultiple2(d, b int) int {
	return b - b % d
}
