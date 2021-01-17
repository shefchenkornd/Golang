package main

import "fmt"

func main() {
	fmt.Println(MaxBall2(37))
}

func round(x, unit float64) float64 {
	return float64(int64(x /unit + 0.5)) * unit
}

func MaxBall2(v0 int) int {
	return int(round((float64(v0)/ 3.6) / 0.981, 1.0))
}