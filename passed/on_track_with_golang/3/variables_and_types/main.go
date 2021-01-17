package main

import (
	"fmt"
	"os"
)

// " := " - автоматически выяснить тип данных и объявить переменную с переданным значением
// или
// var acorn int - статическая типизация помогает Go compiler проверять тип ошибки перед выполнением программы
func main() {
	args := os.Args

	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("hello, Gopher!")
	}
}
