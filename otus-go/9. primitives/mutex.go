package main

import (
	"fmt"
	"sync"
)

var i int // y == 0

func worker(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()   // ВОТ ЭТО
	i++         // КРИТИЧЕСКАЯ СЕКЦИЯ !!!!!!!!!!!!!!!!!!
	mu.Unlock() // поэтому Unlock() должен закрывать эту критическую секцию
	wg.Done()   // а только потом всё остальное!
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &mu)
	}

	wg.Wait()

	// значение переменной "y" будет разное (902, 930 и тд.), но не
	// не будет равным 1000, а почему?
	// потому что у нас доступ к переменной "y" будет одновременно
	// у разных горутин, и они будут переписывать переменную "y"
	fmt.Println("value of y after 1000 operation is", i)
}
