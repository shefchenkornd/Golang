package main

import "fmt"

func main() {
	// преобразоввать строку в слайс байтов или рун
	s := "привет"

	slb := []byte(s)
	slr := []rune(s)

	fmt.Printf("%v \n", slb)
	fmt.Printf("%v \n", slr)
}
