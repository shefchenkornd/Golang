package main

import (
	"fmt"
	"time"
)

// реализация паузы с помощью методов time.Sleep и time.After
func main() {
	fmt.Println("Start....")
	time.Sleep(5 * time.Second)

	fmt.Println("i'm here")

	timeChannel := time.After(12*time.Second)
	<-timeChannel

	fmt.Println("...Finish")
}
