package main

import (
	"fmt"
	"reflect"
	"time"
)

type MyError struct {}

func (e MyError) Error() string {
	return "smth happened"
}

func main() {
	start := time.Now()
	var e error
	e = MyError{}

	/**
	мы можем узнать ТИП структуры либо через рефлексию или через fmt.Printf("%T\n", e)
	 */
	fmt.Println(reflect.TypeOf(e).Name())
	fmt.Printf("%T\n", e)

	fmt.Printf("%v\n", time.Since(start))
}
