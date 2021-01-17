package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	type pair struct {
		i int
		s string
	}

	test := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}

	// где %b - base 2
	fmt.Printf("%b %b %b \n", 0xE2, 0x99, 0xAC)
	fmt.Printf("%d \n", 'А') // 1040
	fmt.Printf("%v \n", []byte("А")) // [208 144]
	fmt.Printf("%b %b \n", 208, 144) // 11010000 10010000 в двоичной системе

	r, s := utf8.DecodeRuneInString("ы")
	fmt.Printf("%v %v \n", r, s) // 11010000 10010000 в двоичной системе

	fmt.Printf("%v \n", utf8.RuneCountInString("Привет")) // 11010000 10010000 в двоичной системе

	var str string
	for i := range test {
		str = atoi(test[i].i)
		if str == test[i].s {
			fmt.Println(test[i].i, str, " - OK!")
		} else {
			fmt.Println(test[i].i, str, " - FAIL!")
		}
	}
}

// integet to acsii
func atoi(i int) (s string) {
	negative := i < 0
	if negative {
		i = 0 - i
	}

	if i == 0 {
		return "0"
	}

	for i > 0 {
		tmp := i % 10
		i = i / 10
		s = string('0'+tmp) + s
	}

	if negative {
		s = "-" + s
	}

	return s
}
