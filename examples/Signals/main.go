package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)


	go func() {
		sig := <- sigs
		fmt.Println() // ^C
		time.Sleep(time.Second * 5)
		fmt.Println(sig) // interrupt
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")



}
