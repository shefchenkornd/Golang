package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Домашнее задание:

Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах)

Функция принимает на вход:
	- слайс с заданиями `[]func()error`
	- число заданий, которые можно выполнить параллельно (`N`)
	- максимальное число ошибок после которого нужно приостановить обработку.

Учесть, что задания могут выполняться разное время!
*/
func t1() int {
	time.Sleep(3 * time.Second)
	fmt.Println("end t1")
	return 0
}

func t2() int {
	time.Sleep(6 * time.Second)
	fmt.Println("end t2")
	return 0
}

func t3() int {
	time.Sleep(8 * time.Second)
	fmt.Println("end t3")
	return 0
}

func t4() int {
	time.Sleep(1 * time.Second)
	fmt.Println("end t4")
	return 0
}

func main() {
	var wg sync.WaitGroup

	N := 3
	finished := make(chan int, N)
	//done := make(chan bool)
	sl := []func() int{
		t1, t2, t3, t4,
	}

	i := 1
	for {
		if i == N {
			// подождать пока освободиться горутина
		}

		l := N
		if 0 > len(sl)-N {
			l = len(sl)
		}

		// Запуск (старт)
		subSl := sl[:l]
		for _, v := range subSl {
			wg.Add(1)
			go func(v func() int) {
				finished <- v()
				wg.Done()
			}(v)
		}
		sl = sl[l:]

		// рабочий цикл, где мы по одному значению вытаскиваем из слайса sl и запускаем в горутине
		select {
		case <-finished:
			fmt.Println(len(sl))
			if len(sl) == 0 {
				fmt.Println("Выходим из select")
				break
			}

			x := sl[0]
			wg.Add(1)
			go func(x func() int) {
				x()
				wg.Done()
			}(x)

			sl = sl[1:]
		}

		if len(sl) == 0 {
			fmt.Println("Выходим из цикла")
			break
		}
	}
	wg.Wait()
}
