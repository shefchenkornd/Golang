package main

import (
	"sync"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		// запуск обработчика Telegram сообщений
		wg.Done()
	}()
	go func() {
		// запуск http сервера
		wg.Done()
	}()
	go func() {
		// запуск ещё какого либо обработчика
		wg.Done()
	}()

	wg.Wait()
}