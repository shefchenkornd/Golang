package main

import "fmt"

func main() {
	fmt.Println(Factorial2(5))
}

func Factorial2(n int) int  {
	if n < 2 {
		return 1
	}

	return n * Factorial2(n-1)
}
