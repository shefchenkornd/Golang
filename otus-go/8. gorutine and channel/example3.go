package main

import "fmt"

/**
когда мы выходим из главной горутины,
выполнение всех остальных горутин тоже прекращается
*/
func main() {
	ch := make(chan int)
	close(ch)
	fmt.Println(<-ch)

	ch <- 10 // panic: send on closed channel
}
