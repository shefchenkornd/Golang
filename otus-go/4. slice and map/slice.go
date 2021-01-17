package main

import "fmt"

func main() {
	var sl []int                         // не-инициализированный слайс
	slic := []string{"a", "b", "c", "d"} // с помощью литерала
	s := make([]int, 3, 10)              // c помощью функции make()

	/**
	функция make() используется для создания:
	* slice
	* map
	* channel
	*/

	fmt.Printf("%T %T %T \n", sl, slic, s)
	fmt.Printf("%v %v %v \n", sl, slic, s)
}
