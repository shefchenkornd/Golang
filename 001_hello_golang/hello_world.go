package _01_hello_golang

import (
	"fmt"
)

func getName() string {
	return "World!"
}

func main() {
	name := getName()
	fmt.Println("Hello "+ name)
}
