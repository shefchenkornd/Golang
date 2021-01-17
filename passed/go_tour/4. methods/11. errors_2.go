package main

import (
	"fmt"
	"reflect"
	"time"
)

type ErrNegativeSqrt struct {
	What string
	When time.Time
}

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %v, %s",
		err.When, err.What)
}

func Sqrt(number float64) (float64, error)  {
	if 0 > number {
		return number, &ErrNegativeSqrt{
			What: "cannot Sqrt negative number: " + fmt.Sprint(number),
			When: time.Now(),
		}
	}

	return 0, nil
}


func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.ValueOf(x))

	v := reflect.ValueOf(x)

	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(111, y)

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-45))
}
