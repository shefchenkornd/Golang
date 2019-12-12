package main

import "fmt"

const ANGLES_ALL = 180

func main() {
	a := 30
	b := 80

	fmt.Println(OtherAngle(a, b))
}

func OtherAngle(a int, b int) int {
	return ANGLES_ALL - (a + b)
}
