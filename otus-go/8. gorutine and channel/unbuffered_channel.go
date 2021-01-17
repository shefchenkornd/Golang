package main

import "fmt"

/**
небуферизованный канал
часто используются для синхронизации
 */
func main() {
	ch := make(chan struct{})

	go func() {
		fmt.Printf("Hello")
		ch <- struct{}{}
	}()

	<-ch
}

