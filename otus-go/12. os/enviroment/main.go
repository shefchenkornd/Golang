package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Environ() // return все env из окружения OS
	fmt.Println(env[10])
	fmt.Printf("\n %v", env)
}
