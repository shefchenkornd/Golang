package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func logger(prefix string) func(s string) {
	return func(s string) {
		fmt.Printf("[%s] %s\n", prefix, s)
	}
}


func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())	// 1
	fmt.Println(nextInt())	// 2
	fmt.Println(nextInt())	// 3


	nextInt2 := intSeq()

	fmt.Println(nextInt2())	// 1


	warn := logger("WARN")
	err := logger("ERROR")
	warn("TEST")
	err("TEST!!!!")
}
