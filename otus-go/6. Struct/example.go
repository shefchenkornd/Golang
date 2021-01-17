package main

import "fmt"

type A struct {
	Field int
}


type B struct {
	A // вот так работает НАСЛЕДОВАНИЕ
	Field int
}

/**
создаем конструктор для структуры А
как пользователь узнает о существовании моего конструктора?
Ответ: почитает документацию

Геттеры и сеттеры в Го обычно не делают, а только по необходимости
 */
func NewA() *A {
	return &A{Field: 12}
}

func size() int {
	return 4 << 10
}

func main() {
	a := A{Field: 3}
	b := a // потому что переменная "b" выделила память и скопировала полностью память переменной "а"

	b.Field++

	d := B{}
	d.A.Field = 27
	d.Field = 34

	fmt.Printf("%v %v \n", a, b)
	fmt.Printf("%v \n", size())
	fmt.Printf("%v %v \n", d.Field, d.A.Field)
}
