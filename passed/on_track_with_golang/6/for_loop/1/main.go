package main

import "fmt"

// в Golang существуе только один цикл - это for {}

// <init> - executed before the first iteration
// <condition> - evaluated before every iteration
// <post> - executed at the end of every iteration

// 	for <init>; <condition>; <post> {}
// или вот так с одним условием <condition>
// 	for <condition> {}
func main() {
	// Declare variables and assign initial values.
	i := 0
	isLessThanFive := true
	for isLessThanFive {
		if i >= 5 {
			// change the condition expression and stops the loop before the next run.
			isLessThanFive = false
		}

		fmt.Println(i)
		i++
	}
}
