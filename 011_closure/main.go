package main

import "fmt"

func main() {

	// анонимная функция
	func(s string) {
		fmt.Println(s)
	}("some word")

	closure()
}

func closure() {
	t := 100
	result := func() {
		fmt.Println(t)
	}

	z := t + 100
	second := func() {
		fmt.Println(z)
	}

	result()
	second()

}
