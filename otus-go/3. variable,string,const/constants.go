package main

import "fmt"

func main() {
	// Константы - неизменяемые значения, доступные только во время компиляции.
	const PI = 3 // принимает подходящий тип
	const pi = 3.14 // строгий тип

	const (
		TheA = "A"
		TheB = "B"
	)

	// const iota = 0 // Untyped int.
	const (
		X = iota // 0
		Y		 // 1
 		Z		 // 2 и тд.
	)

	// 3 int, 3.14 float64, 0, 1, 2
	fmt.Printf("%v %T, %v %T, %v, %v, %v", PI, PI, pi, pi, X, Y, Z)
}
