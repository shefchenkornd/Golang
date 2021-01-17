package main

import (
	"fmt"
	"time"
)

// НЕПРАВИЛЬНОЕ ЗАКРЫТИЕ КАНАЛА
func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	go send(msg)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			close(msg) // здесь мы получим такую ошибку "...panic: send on closed channel"
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

// отправляем сообщение "hello" каждые полсекунды
func send(ch chan string) {
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}
