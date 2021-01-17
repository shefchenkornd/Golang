package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Задача
// запустить функцию через консоль с аргументами и без них
// go run 4.2\ error_example.go hello world

func main() {
	args := os.Args[1:]
	//fmt.Println(args)
	if result, err := Concats(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenatd string: '%s'\n", result)
	}
}


// префикс ... перед типом данных означает, что функция принимает произвольное кол-во аргументов этого типа.
// Язык Golang свернет их в массив этого типа и интерпретирует как []string
func Concats(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no string supplied")
	}

	return  strings.Join(parts, " "), nil
}
