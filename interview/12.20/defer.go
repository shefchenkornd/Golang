package main

import "fmt"

var x int

func f()  int {
	x++
	return x
}

func main() {
	o := fmt.Print

	defer o(f())
	defer func() {
		defer o(recover())
		o(f())
	}()

	defer f()
	defer recover()

	panic(f())
}
