package main

import "fmt"

func main() {
	var str =  "shalom"
	fmt.Printf("%c", str[5]) // Выводит: m

	// строку в Go отредактировать нельзя:
	// str[5] = 'd' // Нельзя присвоить message[5]
	fmt.Println("\n")

	for index := range str  {
		fmt.Printf("%c\n", str[index])
	}
}
