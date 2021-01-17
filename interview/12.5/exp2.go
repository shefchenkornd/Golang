package main

import "fmt"

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1

	return x
}

func main() {
	fmt.Println(test()) // 2
}
