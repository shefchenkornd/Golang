package main

import "fmt"

func main () {
	var num = 23
	// анонимные функции
	multiple := func() int {
		return num * 3
	}

	fmt.Println(multiple())
}