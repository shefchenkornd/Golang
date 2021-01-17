package main

import "fmt"

/**
приведение типов
проверка типа возможно только для ИНТЕРФЕЙСА
*/

/**
x.(T) проверяем, что конкретная часть значения х имеет тип Т и x != nill
- если Т - не интерфейс, то проверяем, что динамический тип х это Т
- если Т - это интерфейс, то проверяем, что динамический тип х его реализует
 */
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	r, ok := i.(fmt.Stringer)
	fmt.Println(r, ok) // <nil> false

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false
}
