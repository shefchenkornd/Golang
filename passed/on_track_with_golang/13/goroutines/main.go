package main

import (
	"fmt"
	"math"
	"sync"
)

// concurrency and parallelism are NOT the same thing.

// Concurrency Allows Parallelism
// The former means independent, which is a necessary step toward the latter, which means simultaneous

// Concurrent means Independent
// Parallel means Simultaneous

// Go's concurrency model and gorounines make it simple to build parallel programs that take advantage of machines with multiple processors.

// ******************************************************************
// if we run our final code specifying a single processor, there's no noticeable performance improvement
//  	выполнить программу на ОДНОМ ПРОЦЕССОРЕ
//		time GOMAXPROCS=1 go run main.go

// The Go runtime defaults to using ALL processors available.
// ******************************************************************
func main() {
	// Go programs are NOT automatically aware of newly created goroutines, so the main function exits before the goroutines are finished.
	// Adding Synchronization with WaitGroup ("sync" package)
	var wg sync.WaitGroup

	names := []string{"Phil", "Bony", "Barbaro"}

	wg.Add(len(names))
	for _, name := range names {
		go printName(name, &wg) // must pass references  to WaitGroup so that we call Done on the original value
	}
	wg.Wait() // prevents program from exiting.
}

func printName(name string, wg *sync.WaitGroup)  {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(name)))
	}

	fmt.Println("Name: ", name)
	wg.Done() // like saying, "Hey, there's one less goroutine you need to wait for."
}