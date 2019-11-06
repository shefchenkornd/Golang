package main

import "fmt"

// такую переменную вне ф-и main() можно объявить и проинициализировать только
// с полным синтаксисом, с краткой инициализацией вызовет ошибку
var outerVariable = "101 долматинец"

func main() {
	var a int = 5
	var b, c string = "b", "c"

	// %T - a Go-syntax representation of the type of the value
	// %v - the value in a default format
	fmt.Printf("%T -  %v\n", a, a)
	fmt.Printf("%T -  %v\n", b, b)
	fmt.Printf("%T -  %v\n\n", c, c)


	// короткий синтаксис
	e := 194
	f := "f"
	g := true

	fmt.Printf("%T -  %v\n", e, e)
	fmt.Printf("%T -  %v\n", f, f)
	fmt.Printf("%T -  %v\n\n", g, g)

	var zeroInt int
	var emptyString string
	var boolFalse bool

	fmt.Printf("%T -  %v\n", zeroInt, zeroInt)
	fmt.Printf("%T -  %v\n", emptyString, emptyString)
	fmt.Printf("%T -  %v\n", boolFalse, boolFalse)


	fmt.Printf("%T -  %v\n", outerVariable, outerVariable)


}
