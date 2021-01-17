package main

import (
	"fmt"
	"strconv"
)

type ErrNegativeSqtr struct {
	Negative float64
}

func (e ErrNegativeSqtr) Error() string {
	return "cannot Sqtr negative number: " + fmt.Sprintf("%v", e.Negative)
}

func main() {
	fmt.Println(Sqtr(2))
	fmt.Println(Sqtr(-2))
}

func Sqtr(x float64) (float64, error) {
	z := 1.0

	if x < 0 {
		return 0, &ErrNegativeSqtr{
			Negative: x,
		}
	}

	for {
		tempResult := formula(z, x)

		koren := tempResult * tempResult
		korenFormatted, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", koren), 64)

		if korenFormatted > x {
			z -= 0.005
		} else if korenFormatted < x {
			z += 0.005
		} else if korenFormatted == x {
			return tempResult, nil
		}
	}
}

func formula(z, x float64) float64 {
	return z - (z*z-x)/2*z
}