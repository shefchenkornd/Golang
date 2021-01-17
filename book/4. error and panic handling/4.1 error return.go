package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Concat())


	parts := "Error string should not be capitalized or end with punctuation "
	fmt.Println(Concat(parts, "two arg", "three "))
}

// префикс ... перед типом данных означает, что функция принимает произвольное кол-во аргументов этого типа.
// Язык Golang свернет их в массив этого типа и интерпретирует как []string
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no string supplied")
	}

	return  strings.Join(parts, " "), nil
}