package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")

	say("hello")

	// чтобы go-routine не завершалась, либо ставим ф-ю которая ждёт пользовательский ввод или ....
	//fmt.Scanln()

	fmt.Println("main finished")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}
