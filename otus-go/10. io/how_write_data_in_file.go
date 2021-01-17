package main

import (
	"io"
	"log"
	"os"
)

/**
в отличии от чтения тут цикл не нужен
 */
func main() {
	b := make([]byte, 1024*1024)

	file, _ := os.Create(path)
	defer file.Close() // чтобы очистить буфер ОС

	file.Seek(io.SeekCurrent, 0)
	written, err := file.Write(b)
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}

	// мы записали 1Mb данных !
}
