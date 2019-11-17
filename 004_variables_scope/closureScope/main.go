package main

import "fmt"

func main() {
	x := 100
	fmt.Println(x)
	someFunc()

	fmt.Println(miniGlobal)
}

func someFunc() {
	x := 200
	fmt.Println(x)
}

var miniGlobal = "miniGlobal"