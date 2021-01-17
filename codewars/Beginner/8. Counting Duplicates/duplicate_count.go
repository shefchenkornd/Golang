package main

import (
	"fmt"
	"strings"
)

// Count the number of Duplicates
// Write a function that will return the count of distinct case-insensitive alphabetic characters and
// numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets
// (both uppercase and lowercase) and numeric digits.
func main() {
	fmt.Println(duplicate_count("indivisibility"))
}

func duplicate_count(s1 string) (result int) {
	s1 = strings.ToLower(s1)

	m := make(map[rune]rune)

	for _, rune := range s1 {
		if strings.Count(s1, string(rune)) > 1 {
			if _, ok := m[rune]; !ok {
				m[rune] = rune
				result++
			}
		}
	}

	return
}