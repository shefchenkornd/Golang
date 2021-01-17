package main

import (
	"fmt"
)

func loop(n int) {
	for i := 1; i < 10000; i++ {
		fmt.Printf("%v", n)
	}
}

func main() {
	// Асинхронность в Го
	//runtime.GOMAXPROCS(1)

 	go loop(1)
 	go loop(2)
	fmt.Scanln()
}
