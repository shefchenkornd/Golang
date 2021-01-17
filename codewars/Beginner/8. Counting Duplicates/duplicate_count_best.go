package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(duplicate_count2("indivisibility"))
}

func duplicate_count2(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}