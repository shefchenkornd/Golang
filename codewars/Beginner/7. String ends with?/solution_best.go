package main

import (
	"fmt"
	"strings"
)

func main() {
	str, ending := "abc", "c"
	//str, ending := "", ""
	//str, ending := "    ", ""
	//str, ending := "    ", "    "
	//str, ending := "100", "1oo"
	fmt.Println(solution2(str, ending))
}

// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).
func solution2(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
