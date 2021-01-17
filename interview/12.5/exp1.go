package main

import "fmt"

type T struct {
	i int
	f float64
	next *T
}

func main() {
	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := numbers[:5:8]
	fmt.Println(s1)

	fmt.Println(T{}) // {0 0 <nil>}
}
