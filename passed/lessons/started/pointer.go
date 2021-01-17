package main

import "fmt"

func main () {
	var x  = 0

	pointer (&x)

	fmt.Println(x)
}

func pointer(x *int) {
	*x = 2
}
