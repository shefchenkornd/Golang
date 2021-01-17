package main

import (
	"bufio"
	"os"
)

/**
для чтения/записи
чтобы повысить производительсть вашего приложения
нужно читать не кусками, а читать большой файл в буфер (bufio),
затем считывать из буфера как вам удобно и записывать в файл, к примеру!
 */
func main() {
	path := "/path"
	file, _ := os.Create(path)
	bw := bufio.NewWriter(file)
	writter, err := bw.Write([]byte("some bytes"))
	bw.WriteString("some string")
	bw.WriteRune('+')
	bw.WriteByte(42)
	bw.Flush() // очистить буфер, записать всё в file

}
