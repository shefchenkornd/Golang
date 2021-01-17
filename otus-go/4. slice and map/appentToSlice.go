package main

import "fmt"

func main() {
	var s []int // zero-value для slice  будет nil
	s3 := []int{1, 2, 3, 4, 5}
	s = append(s, s3...) // разложили слайс по item'ам

	/**
	нативной функции удаления в slice нет, по сути придется создавать новый слайс из частей оригинального
	 */

	fmt.Printf("%v", s)
}
