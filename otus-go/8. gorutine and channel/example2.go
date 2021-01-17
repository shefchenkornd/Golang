package main

import "fmt"

/**
когда мы выходим из главной горутины,
выполнение всех остальных горутин тоже прекращается
*/
func main() {
	//go fmt.Printf("Hello") // ничего не выведет!
	// для вывода "Hello" на нужны каналы

	ch := make(chan bool)
	go func() {
		ch <- true
	}()
	fmt.Println(<-ch)
}
