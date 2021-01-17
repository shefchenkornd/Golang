package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "hello"
	}()

	timer := time.NewTimer(2 * time.Second)

	select {
	case data := <-ch:
		fmt.Printf("Received %v", data)
	case <-timer.C:
		fmt.Println("End!!!")
	}
}
