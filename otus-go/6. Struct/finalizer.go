package main

import (
	"fmt"
	"runtime"
	"time"
)

type Test struct {
	A int
}

type T struct {
	_ int
	_ bool
}

func test() {
	// создаем указатель
	a := &Test{}
	// добавляем простой финализатор
	runtime.SetFinalizer(a, func(a *Test) {
		fmt.Println("I am dead")
	})
}

func main() {

	var t2 = T{12,false}
	println(t2)

	test()
	// запускаем сборку мусора
	runtime.GC()
	// даём время горутине финализатора отработать
	time.Sleep(1 * time.Millisecond)
}
