package main

import (
	"fmt"
	"sync"
	"time"
)

type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk() {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
}

func main() {
	dogs := []Dog{
		{"Bobik", time.Second},
		{"Tuzik", time.Second * 3},
	}

	// взяли указатель, потому что если WaitGroup передаём куда-то дальше,
	// чтобы эта была одна и та же WaitGroup!
	wg := &sync.WaitGroup{}

	// нет гарантий, что эта функция что-то выведет!
	// если мы не используем WaitGroup - механизм ддя ожидания завершения работы
	// нескольких горутин!
	for _, d := range dogs {
		wg.Add(1)

		// ВНИМАНИЕ: нужно прокинуть переменную "d"
		// внутрь анонимной функции, чтобы работало корректно!
		go func(d Dog) {
			defer wg.Done()
			d.Walk()
		}(d)
	}
	wg.Wait()
	fmt.Println("Everybody's home")
}
