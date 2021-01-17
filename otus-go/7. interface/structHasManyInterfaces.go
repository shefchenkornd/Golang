package main

import "fmt"

/**
Одна структура может реализовать несколько интерфейсов
 */
type Hound interface {
	hunt()
}

type Poodle interface {
	bark()
}

type Rottweiler struct {
	name string
}

func (r Rottweiler) hunt()  {
	fmt.Println("hunt")
}

func (r Rottweiler) bark()  {
	fmt.Println("bark")
}

/**
работаем через интерфейс
 */
func f1(i Hound) {
	i.hunt()
}

/**
работаем через интерфейс
*/
func f2(i Poodle) {
	i.bark()
}

func main() {
	rottweiler := Rottweiler{name: "poodle"}
	rottweiler.hunt() // hunt
	rottweiler.bark() // bark

	f1(rottweiler) // hunt
	f2(rottweiler) // bark
}
