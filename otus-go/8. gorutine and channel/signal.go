package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
	скомпилировать и запустить бинарник
	отправить бинарнику сигналы 2 и 15, как на примере ниже
	kill -2 PID
	kill -15 PID
 */
func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	fmt.Printf("Got %v...\n", <-interrupt)
}
