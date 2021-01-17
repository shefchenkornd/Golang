package main

import "fmt"

/**
Одному интерфейсу могут соответствовать много типов
 */

type Poodle2 interface {
	Barks()
}

type Clip struct {
	name string
}

func (Clip) Barks() {
	fmt.Println("Barks")
}

type ToyPoodle struct {
	name string
}

func (ToyPoodle) Barks() {
	fmt.Println("Barks")
}

func main() {
	var clip, toypoodle Poodle2

	clip = Clip{name: "Clip"}
	toypoodle = ToyPoodle{name: "ToyPoodle"}

	clip.Barks()      // Barks
	toypoodle.Barks() // Barks
}
