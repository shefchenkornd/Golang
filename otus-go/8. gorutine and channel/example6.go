package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Print(i) // выведет: 03412 , НО порядок не будет гарантирован!!!
		}(i)
	}

	time.Sleep(10 * time.Second)
}
