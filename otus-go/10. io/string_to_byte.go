package main

import "fmt"

func main() {
	s := "some string"
	b := []byte(s)

	fmt.Printf("%p %p", &s, &b)
}
