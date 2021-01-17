package main

import (
	"fmt"
	"time"
)

// использование завершающего канала
func main() {
	msg := make(chan string)
	done := make(chan bool) // допол.канал  с типом данных bool для сообщения о завершении
	until := time.After(5 * time.Second)

	go send2(msg, done)
	for {
		select {
		case m:=<-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send2(ch chan<-string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch<-"hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}