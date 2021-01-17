package main

import "fmt"

func main() {
	var a [4]int                 // initialize zero-value
	b := [4]int{1, 2, 3, 4}      // используем литералы
	c := [...]int{5, 6, 7, 8, 9} // [...] - длина массива будет вычислена автоматически
	multiArr := [5][5]int{}      // многомерный массив

	fmt.Printf("%v %v %v %v \n", a, b, c, multiArr)

	/*
	ОПЕРАЦИИ НАД МАССИВОМ
	c[3] = 1 // indexing
	len(c)	 // array length
	с[2:3]	 // getting slice
	 */

	fmt.Printf("Indexing = %v,\n length = %v,\n Slice = %v,\n", c[3], len(c), c[2:])
}
