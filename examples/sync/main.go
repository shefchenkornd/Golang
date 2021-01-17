package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan bool)

	id := os.Getpid()
	fmt.Println(id)

	go func() {
		fmt.Println("Hello....")

		/**
		после закрытия вы всегда будете получать zero value и чтение из кнала закончено
		*/
		close(ch) // or ch <- true
	}()

	<-ch
}
