package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string)  {
	c.mux.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++

	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.v[key]
}

func main() {
	safecounter := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go safecounter.Inc("somekey")
	}


	time.Sleep(time.Second)
	fmt.Println(safecounter.Value("somekey"))
}
