package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	fmt.Println(len(question)) // String: the number of bytes in v.
	fmt.Println(utf8.RuneCountInString(question)) // вернёт кол-во символов в строке

	k := "¿"
	firstCharacter, size := utf8.DecodeRuneInString(k)
	fmt.Printf("First rune: %c %v bytes\n", firstCharacter, size) // Функция DecodeRuneInString возвращает первый символ и количество байт, что символ тратит

	for i, c := range question {
		fmt.Printf("rune - %v | character - %c  \n", i, c)
		//fmt.Printf("%c ", c)
	}
}