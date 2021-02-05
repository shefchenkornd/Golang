package main

import "fmt"

/**
форматированный вывод
методы начинающиеся на `Fprint...` ждут первым параметром io.Writer !!!


Основные флаги
общие:
	%v - представление по умолчанию для типа
	%#v - вывести как Go код (удобно для структур)
	%T - вывести тип переменной
	%% - вывести символ `%`

для целых:
	%b	base 2
	%d	base 10
	%o	base 8
	%x	base 16, with lower-case letters for a-f

для строк:
	%s	как есть
	%q 	вывести строку в двойных кавычках
	%x 	в 16ом формате

pointer:
	%p - pointer (адреса переменных)
 */
func main() {
	fmt.Printf("%s \n", "abs")	// abs
	fmt.Printf("%q \n", "abs")	// "abs"
	fmt.Printf("%x \n", "abs")	// 616273

	fmt.Printf("%#v", struct {
		name string
		family string
		age int
	}{})
	// output: struct { name string; family string; age int }{name:"", family:"", age:0}
}