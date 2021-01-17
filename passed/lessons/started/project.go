package main

import (
	"fmt"
	"unicode/utf8"
)

func main () {
	//const PI = 3.14
	var age = 15
	var pi float32 = 2.143
	var str = "Hello world!"
	var str_ru = "привет!"

	var a = 3
	var b = 5
	var res int
	res = a + b
	fmt.Println(age, pi,str, res)

	var s int
	s = utf8.RuneCountInString(str_ru)
	fmt.Println(len(str), s)


	fmt.Println(str + " cool \nis Fun" )

	if  age < 5 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	switch age {
	case 5: fmt.Println("Вам 5 лет")
	case 10: fmt.Println("Вам 10 лет")
	case 15: fmt.Println("Вам 15 лет")
	default:
		fmt.Println("Сколько вам лет?")
	}

	var i = 1
	for i <=10  {
		fmt.Println(i)
		i++
	}




	var arr[3] int
	arr[0] = 3
	arr[1] = 75
	arr[2] = 9

	fmt.Println("Выводим массив", arr)


	nums := [3]float64 {3.23, 45.22, 93.3}

	for i, value := range nums {
		fmt.Println(i, value)
	}

	webSites := make(map[string]float64)

	webSites["itProger"] = 0.8
	webSites["yandex"] = 0.99

	fmt.Println("Выводим maps", webSites["yandex"])

	arg_1, arg_2 := summ (10, 4)
	fmt.Println("Результат функции summ() =", arg_1, arg_2)
}
 //godoc fmt Println

func summ (num_1 int, num_2 int) (int, int)  {
	var res int
	res = num_1 + num_2
	return res, num_1 * num_2
}