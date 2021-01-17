package main

import "fmt"

func main() {
	/**
	Аргументы отложенного вызова ф-ции вычисляются тогда, когда вычисляется команда defer
	(т.е. аргументы рассчитываются до выполнения defer)
	 */
	i:=0
	defer fmt.Println(i) // output: 0
	i++
	fmt.Println("i = ", i) // output: 1
}
