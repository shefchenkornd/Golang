package main

import "fmt"

/*
Каналы можно сравнить!
Каналы равны только в том случае, если у них одинаковые указатели
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := ch2

	fmt.Println("ch2 == ch3", ch2 == ch3)
	fmt.Println("ch1 == ch2", ch1 == ch2)
	fmt.Println(ch1, ch2, ch3) // 0xc00001e180 0xc00001e1e0 0xc00001e1e0
}
