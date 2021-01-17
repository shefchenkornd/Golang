package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str, ending := "1oo", "100"
	fmt.Println(solution(str, ending))
}

// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).
func solution(str, ending string) bool {
	var index int = utf8.RuneCountInString(ending)

	if index == 0 { // означает пустая строка
		return true
	}

	for i, _ := range str  {
		if str[i:] == ending {
			return true
		}
	}

	return false
}
