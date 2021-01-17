package main

import (
	"io"
	"log"
	"os"
)

func main() {
	N := 1024              // мы знаем сколько хотим прочитать
	buf := make([]byte, N) // подготавливаем буфер нужного размера

	file, _ := os.Open(path) // открываем файл
	offset := 0

	for offset < N {
		read, err := file.Read(buf[offset:])
		offset += read
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
	}
	// мы прочитали N байт в buf !
}
