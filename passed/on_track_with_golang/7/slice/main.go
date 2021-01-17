package main

import "fmt"

func main() {
	var langs []string

	// добавлять в slice можно только так, способ добавления элемента в массив, как в PHP здесь не работает
	langs = append(langs, "Go!")
	langs = append(langs, "Ruby")
	langs = append(langs, "Js")
	langs = append(langs, "LoL")

	fmt.Println(langs)
}
