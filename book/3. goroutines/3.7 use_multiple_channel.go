package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(30 * time.Second) // это канал, посылающий сообщение по истечении 30 секунд.
	echo := make(chan []byte)            //создание нового канала для передачи байтов из Stdin в Stdout. Поскольку размер не указан, то канал может хранить только одно сообщение одновременно

	go readStdin(echo) // запуск горутины для чтения данных из Stdin и перадачи их в новый канал

	for { // используем select для передачи данных из Stdin в Stdout, если  имеются и для завершения по событию таймера
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

// принимает канал для записи (chan<-) и отправляет в канал введенные данные
func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024) // копирование данных из Stdin в объект data.
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data // отправка буферизированных данных в канал
		}
	}
}
