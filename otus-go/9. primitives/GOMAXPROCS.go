package main

import (
	"fmt"
	"runtime"
	"sync"
)

var y int // y == 0

func worker2(wg *sync.WaitGroup, mu *sync.Mutex) {
	//mu.Lock()
	//mu.Unlock()
	y++
	wg.Done()
}

func main() {
	/**
	если GOMAXPROCS выставить в единицу, Mutex нам не потребутся, но это
	больше лайфхак, чем решение проблемы!
	 */
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker2(&wg, &mu)
	}

	wg.Wait()

	// значение переменной "y" будет разное (902, 930 и тд.), но не
	// не будет равным 1000, а почему?
	// потому что у нас доступ к переменной "y" будет одновременно
	// у разных горутин, и они будут переписывать переменную "y"
	fmt.Println("value of y after 1000 operation is", y)
}
