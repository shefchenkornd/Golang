package main

import (
	"fmt"
	"io"
)

// пустой интерфейс
type Empty interface{}

func Fprintln(w io.Writer, a ...interface{}) {
	// do it
}

func printAll(values []interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main()  {
	names := []string{"Джеки", "Коля", "Степан"}

	// Can I convert a []T to an []interface{}? ¶
	// Напрямую так сделать нельзя, только через копирование одного слайса в другой с нужным типом
	s := make([]interface{}, len(names))
	for i, v := range names {
		s[i] = v
	}

	//fmt.Printf("%v", s)


	printAll(s)
}
