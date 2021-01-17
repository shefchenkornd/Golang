package main

import "fmt"

func main() {
	// Однако также для получения ввода с консоли мы можем использовать
	// встроенные функции fmt.Scan(), fmt.Scanln() и fmt.Scanf(),
	// которые аналогичны соответственно функциям fmt.Fscan(), fmt.Scanln() и fmt.Scanf():

	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}
