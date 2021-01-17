package main

import "fmt"

func main() {
	a := [][]int{{30,40}, {69, 1300}, {1, 4}}

	fmt.Println(ConvertFracts2(a))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func ConvertFracts2(a [][]int) string {
	d := 1
	for _, num := range a {
		g := gcd(num[0], num[1])
		num[0] /= g
		num[1] /= g
		d = d * num[1] / gcd(d, num[1])
	}
	res := ""
	for _, num := range a {
		res += fmt.Sprintf("(%d,%d)", num[0] * d / num[1], d)
	}
	return res
}