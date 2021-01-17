package main

import "time"

// закрытие отправителем
func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go sent(ch)
	for {
		select {
		case <-ch:
			println("Got message.")
		case <-timeout:
			println("Time out.")
			return
		default:
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func sent(ch chan bool)  {
	time.Sleep(500 * time.Millisecond)
	ch <-true
	close(ch)
	println("Sent and closed.")
}