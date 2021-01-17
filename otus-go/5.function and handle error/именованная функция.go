package main

import "fmt"

func Greet() int {
	return 100
}
func main() {
	fmt.Println( Greet() )

	// или функцию можно объявить так, но через анонимную ф-ю:
	greet := func() int {
		return 200
	}

	fmt.Println( greet() )
}