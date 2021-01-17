package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
Буферизированный ввод-вывод
Большинство встроенных операций ввода-вывода не используют буфер. Это может иметь отрицательный эффект для производительности приложения.
Для буферизации потоков чтения и записи в Go определены ряд возможностей, которые сосредоточены в пакете bufio.
 */
func main() {
	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}

	file, err := os.Create("some.dat") // создаем файл
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(file)
	defer file.Close()	// закрываем файл

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}

	// В данном случае в файл через буферизированный поток вывода записываются две строки.
	writer.Flush()
}
