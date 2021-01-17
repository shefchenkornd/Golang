package main

import (
	"fmt"
	"log"
)

/**
Recover - это встроенная функция, которая восстанавливает контроль над паникующей го-процедурой.
Recover полезна только внутри отложенного вызова функции. Во время нормального выполнения, recover возвращает nil
и не имеет других эффектов. Если же текущая го-процедура паникует, то вызов recover возвращает
значение, которое было передано panic и восстанавливает нормальное выполнение.
 */

/**
"поймать" панику можно с помощью recover: вызов recover останавливает выполнение отложенных функций
и возвращает аргумент, переданный panic
 */

type Work struct {}

func do(w *Work)  {
	if w == nil {
		panic("Panic in do func")
	}
	fmt.Printf("WORK IS DONE \n")
}

func server(workChan <-chan  *Work)  {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work)  {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		} else {
			log.Println("work ok", err)
		}
	}()
	do(work)
}

func main() {
	safelyDo(nil)
	fmt.Println("YEAP!")
}
