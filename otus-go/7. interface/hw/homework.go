package main

import (
	"fmt"
)

/**
interface
*/
type Hound interface { // гончая
	Hunt()
}

type Poodle interface { // пудель
	Bark() string
}

type Decorative interface { // декоративная
	Sleep(hours int)
}

/**
struct
*/
type Bloodhound struct { // ищейка
	name string
}

func (b Bloodhound) Hunt() {
	fmt.Printf("%s is hunting\n", b.name)
}

type TeacupPoodle struct{ name string }

func (p TeacupPoodle) Bark() string {
	return fmt.Sprintf("%s is barking...", p.name)
}

type ShihTzu struct{ name string }

func (p ShihTzu) Sleep(hours int) {
	fmt.Printf("%s is sleeping...\n", p.name)
}

/**
Задание:
 в зависимости от типа собаки выполнить, то что она умеет!
*/
func zoo(dog interface{}) {
	switch dog.(type) {
	case Hound:
		dog.(Hound).Hunt()
	case Poodle:
		fmt.Println(dog.(Poodle).Bark())
	case Decorative:
		dog.(Decorative).Sleep(4)
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	zoo(Bloodhound{"Тузик"})
	zoo(TeacupPoodle{"Бобик"})
	zoo(ShihTzu{"Киша"})
	zoo(int(12))
}
