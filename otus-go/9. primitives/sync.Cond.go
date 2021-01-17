package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Cond
Cond(itional variable) -  механизм для ожидания горутинами сигнала о событии

type Cond struct {
	L Locker
}

func (c *Cond) Broadcast() - будит все горутины, которые ждут `c`.
func (c *Cond) Signal() - будит одну горутину, которая ждёт `c`, если такая есть!
func (c *Cond) Wait() - разблокирует c.L, ждёт сигнала и снова блокирует c.L
 */
type Dogs struct {
	name string
}

func (d *Dogs) Eat(food *DogFood)  {
	food.Lock()
	food.cond.Wait()
	food.food--
	food.Unlock()
}

type DogFood struct {
	sync.Mutex
	food int
	cond *sync.Cond
}

func NewDogFood(food int) *DogFood {
	r := DogFood{food: food}
	r.cond = sync.NewCond(&r)
	return &r
}

func main() {
	var wg sync.WaitGroup

	food := NewDogFood(4)

	for _, d := range []Dogs{{name: "Vasya"}, {name: "Bob"}} {
		wg.Add(1)
		go func(d Dogs) {
			defer wg.Done()
			d.Eat(food)
		}(d)
	}

	println("Waiting for food to arrive....")
	time.Sleep(1 * time.Second)

	food.cond.Broadcast() // мы всех разбудили, кто висит на Broadcast()

	wg.Wait()
	fmt.Printf("Food left: %d\n", food.food)
}
