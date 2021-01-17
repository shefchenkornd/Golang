package main

import "io"

/**
Здесь `dst` должен реализовать интерфейс io.Writer
`src` - io.Reader
 */
func main() {
	// копирует всё вплоть до io.EOF
	written, err := io.Copy(dst, src)

	// копирует N байт или до io.EOF
	written, err := io.CopyN(dst, src, 42)

	// копирует всё вплоть до io.EOF, но использует заданный буфер
	buffer := make([]byte, 1024*1024)
	written, err := io.CopyBuffer(dst, src, buffer)
}
