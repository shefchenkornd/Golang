package main

import (
	"fmt"
	"time"
)

// Задача
// простая блокировка по средствам каналов
func main() {
	lock := make(chan bool, 1) // создаем буферизированный канал единичного объема
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}

	time.Sleep(10 * time.Second)
}

func worker(i int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", i)
	lock <- true // устанавливаем блокировку с помощью отправки сообщения в буфер.канал

	// ############### этот фрагмент кода между [lock <- true/ <-lock] выполняется под защитой блокировки #####################
	fmt.Printf("%d has the lock\n", i)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", i)
	// ##########################################

	<-lock // снять блокировку, прочитав сообщение из канала. В результате в буфере освободится место, и след.ф-я сможет установить блокировку
}
