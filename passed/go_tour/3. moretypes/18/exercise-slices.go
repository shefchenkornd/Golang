package main

import "golang.org/x/tour/pic"

//  программа покажет изображение, интерпретируя целые числа (uint8 они же ака byte) как оттенки серого (вообще-то, синего).
func Pic(dx, dy  int) [][]uint8  {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = make([]uint8, dx)
	}

	for y := range result {
		for x := range result[y] {
			result[y][x] = uint8(x*y)
		}
	}

	return  result
}

func main() {
	pic.Show(Pic)
}
