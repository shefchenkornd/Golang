package main

import "fmt"

func main() {
	// функцию можно запихнуть в переменную
	f:= func() int {
		return 100
	}

	fmt.Println(f())
}
