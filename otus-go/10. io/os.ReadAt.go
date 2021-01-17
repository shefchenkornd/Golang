package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 100)
	f, err := os.Open("/dev/random")
	fmt.Println(err)

	d, err := f.ReadAt(b, 10)
	fmt.Printf("d: %v \nerr: %v \nb: %v \n", d, err, b)

	s := string(b)
	fmt.Printf("string from byte: %v\n", s)
}
