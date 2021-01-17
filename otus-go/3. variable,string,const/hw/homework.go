package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
------------ Решение -----------
защита от дурака
если в строке нет букв (латинских), то отдаем пустую строку
если нет цифр в строке, то отдаём ту же строку

после буквы может быть цифра
а может и нет цифр
то мы пишем в результирующий массив
 */


func main() {
	str := "a4bc2d5e"
	//str := "abcd"
	//str := "45"

	fmt.Println(unpackString(str))
}

/**
 * Распаковка строки
 */
func unpackString(str string) string {
	var result []string

	letterCheck, _ := regexp.MatchString("[a-zA-z]+", str)
	// если в строке нет букв (латинских), то отдаем пустую строку
	if !letterCheck {
		return ""
	}

	digitCheck, _ := regexp.MatchString("[0-9]+", str)
	// если нет цифр в строке
	if !digitCheck {
		return str
	}

	letterSlice := strings.Split(str, "")
	for _, letter := range letterSlice {
		number, err := strconv.Atoi(letter)
		if err == nil {
			lastElem := result[len(result)-1]
			for i := 1; i < number; i++ {
				result = append(result, lastElem)
			}
		} else {
			result = append(result, letter)
		}
	}

	return strings.Join(result, "")
}