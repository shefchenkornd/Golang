package main

import (
	"fmt"
	"runtime"
)

/**
runtime.NumGoroutine() - покажет кол-во запущенных горутин
*/
func main() {
	fmt.Printf(
		"Goroutines: %d",
		runtime.NumGoroutine(),
	)
}
