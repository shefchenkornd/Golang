package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Задание: реализуйте WordCount.
// Она должна вернуть карту с количеством повторений каждого "слова" в строке s.
// Функция wc.Test запускает набор тестов для предоставленной функции и выводит сообщение об успехе или неудаче.

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)

	sl := strings.Fields(s)
	for _, v := range sl {
		if count, ok := wordCount[v]; ok {
			wordCount[v] = count + 1
		} else {
			wordCount[v] = 1
		}
	}

	return wordCount
	//return map[string]int{"x": 1}
}

func main() {
	//fmt.Println(WordCount("The maximum sum subarray problem consists in finding the maximum sum of a contiguous subsequence in an array or list of integers"))
	wc.Test(WordCount)
}