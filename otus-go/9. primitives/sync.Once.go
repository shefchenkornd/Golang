package main

import "sync"

type oldDog struct {
	name string
	die  sync.Once
}

func (d *oldDog) Die() {
	// Метод Do вызывает функцию f только если Do был вызван впервые для этого инстанса Once.
	d.die.Do(func() {
		println("bye!")
	})
}

/**
sync.Once и его метод Do
Once это объект, который выполнит только однократное действие.
Метод Do вызывает функцию f только если Do был вызван впервые для этого инстанса Once.
 */
func main() {
	ch := make(chan struct{})
	close(ch)

	d:= oldDog{
		name: "bob",
	}
	d.Die()
	d.Die()
	d.Die()
	d.Die()
}
