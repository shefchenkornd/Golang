package main

import (
	"fmt"
	"regexp"
	"strings"
)

var longSentece = "Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:"

func reverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func main() {
	// разбить строку на слайс слов
	// очистить слова от знаков пунктуации
	// сделать подсчёт слов

	s := strings.Split(longSentece, " ")
	for k, v := range s {
		// в нижнем регистре
		tpm := strings.ToLower(v)

		// дебильный синтаксис у регул.выражений, но что есть так и живем
		// удаляем символы пунктуации
		re := regexp.MustCompile(`[[:punct:]]`)
		tpm = re.ReplaceAllString(tpm, "")

		s[k] = tpm
	}

	// сделать подсчёт слов
	wordCounter := map[string]int{}
	for _, word := range s {
		if counter, ok := wordCounter[word]; ok {
			counter = counter + 1
			wordCounter[word] = counter
			continue
		}

		wordCounter[word] = 1
	}

	m := map[int]string{}
	m = reverseMap(wordCounter)

	fmt.Println(m)
}
