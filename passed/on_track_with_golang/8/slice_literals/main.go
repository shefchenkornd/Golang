package main

import "fmt"

func main() {
	// slice literals is quick way to create slice with elements
	langs := []string{"Kotlin", "Node.js", "Agail"}

	// получить доступ к элементам слайса можно по индексу или с помощью ключевого слово `range`
	fmt.Println(langs[0])
	fmt.Println(langs[1])
	fmt.Println(langs[2])

	// если нам не известно кол-во элементов в слайсе, то используем ключевого слово `range`
	for indx := range langs {
		fmt.Println(langs[indx])
	}
}
