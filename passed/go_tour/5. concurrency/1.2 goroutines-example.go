package main

import (
	"fmt"
	"sync"
)

var swg sync.WaitGroup // 1

func routine(i int) {
	defer swg.Done() // 3 -- decrement items to be waited by 1
	fmt.Printf("routine %v finished\n", i)
}

func main() {
	swg.Add(10) // 2 -- add 10 items to wait
	for i := 0; i < 10; i++ {
		go routine(i) // *
	}
	swg.Wait() // 4
	fmt.Println("main finished")
}