package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// эхо-вывод в фоновом режиме, пока в основной программе работает таймер
func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(10* time.Second)

	fmt.Println("Timed out.", time.Now().Format("2006-01-01 15:04:05.0000"))
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer)  {
	io.Copy(out, in)
}
//2020-04-11 19:43:33.8561
//2020-04-11 19:43:47.0964