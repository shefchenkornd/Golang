package main

import "fmt"

func main() {
	s := "итерация по строке"

	// по байтам
	for i:=0; i < len(s); i++ {
		b := s[i]
		// i строго последовательно
		// b имеет тип byte, uint8
		fmt.Printf("byte - %v \n", b)
	}

	// по рунам
	for i, r := range s {
		// i может перепрыгивать значения 1,2,4,6,9...
		// r - имеет тип rune, int32
		fmt.Printf("i - %v, rune - %v \n", i, r)
	}
}
