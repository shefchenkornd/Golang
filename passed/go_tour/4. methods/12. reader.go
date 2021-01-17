package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	fmt.Println("Только что созданный срез ", b, "\n")

	for {
		// Read записывает данные в предоставленный срез байтов и возвращает количество записанных байтов и значение типа error.
		// Он возвращает ошибку io.EOF, когда поток заканчивается.
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
