package main

import "fmt"

func main () {
	// отложенная функция
	defer two()

	one()
}

func one()  {
	fmt.Println("1")
}

func two()  {
	fmt.Println("2")
}