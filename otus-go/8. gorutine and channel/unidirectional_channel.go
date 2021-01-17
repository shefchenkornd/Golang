package main

import (
	"fmt"
	"time"
)

func fout(out chan<- int)  {
	out <- 5
}

func fin(in <-chan int)  {
	fmt.Printf("%d\n", <-in)
}

func main() {
	var ch = make(chan int)
	go fout(ch)
	go fin(ch)

	time.Sleep(5 * time.Second)
}
