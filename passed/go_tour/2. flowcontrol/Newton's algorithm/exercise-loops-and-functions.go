package main

import (
	"fmt"
	"strconv"
)

// реализуйте алгоритм нахождения квадратного корня методом Ньютона.
func main() {
	fmt.Println("Квадратный корень из 2 равен: ", Sqrt(2))
}

func Sqrt(x float64) float64 {
	z := 1.0

	for {
		tempResult := formula(z, x)

		koren := tempResult * tempResult
		korenFormatted, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", koren), 64)

		if korenFormatted > x {
			z -= 0.005
		} else if korenFormatted < x {
			z += 0.005
		} else if korenFormatted == x {
			return tempResult
		}
	}
}

func formula(z, x float64) float64 {
	return z - (z*z-x)/2*z
}