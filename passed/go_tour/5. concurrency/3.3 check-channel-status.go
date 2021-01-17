package main

import "fmt"

func main() {
	intCh := make(chan int, 3)

	intCh <- 1
	intCh <- 2

	close(intCh)

	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened {
			fmt.Println(val, opened)
		} else {
			fmt.Println("Channel closed!")
		}
	}
}
