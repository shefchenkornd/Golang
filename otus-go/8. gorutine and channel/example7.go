package main

import "fmt"

const N = 3

func main() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		fmt.Println(&i)
		m[i] = &i // здесь будет ссылка на область памяти, где хранится переменная i
	}

	fmt.Println("********")
	for _, v := range m {
		// здесь мы получаем значение по указателю,
		// а т.к. в переменная i сейчас лежит значение три,
		// то мы получаем вывод: 333
		fmt.Println(*v) // 333
	}

}
