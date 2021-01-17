package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Алаверды"
	subStr := "верды"
	prefix := "Ала"
	suffix := "рды"

	// проверка наличия подстроки
	fmt.Printf("Result Contains %t \n", strings.Contains(str, subStr))

	// строка начинается с ?
	fmt.Printf("Result hasPrefix %t \n", strings.HasPrefix(str, prefix))

	// строка заканчивается на ?
	fmt.Printf("Result hasSuffix %t \n", strings.HasSuffix(str, suffix))

	// склейка строк
	fmt.Printf("Result Join %s \n", strings.Join([]string{"я","ты","он","она","вместе дружная семья"}, ","))

	// разбиение по разделителю
	fmt.Printf("Result split %s \n", strings.Split(str, ""))
}
