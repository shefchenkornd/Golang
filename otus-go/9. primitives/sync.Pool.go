package main

import (
	"fmt"
	"sync"
)

type Dog2 struct {
	name string
}

func (d *Dog2) Bark()  {
	fmt.Printf("%s", d.name)
}

var dogPack = sync.Pool{
	New: func() interface{} {
		return &Dog2{}
	},
}

/**
Object Pool это порождающий паттерн проектирования, используемый для подготовки и
хранения множественных экземпляров согласно заданным параметрам.

Паттерн пула объектов полезен в случаях, когда инициализация объекта более дороже, чем обслуживание объекта.

Паттерн пул объектов положительно влияет на ПРОИЗВОДИТЕЛЬНОСТЬ из-за предварительной инициализации объектов.
*/
func main() {
	dog := dogPack.Get().(*Dog2)
	dog.name = "Tuzik"
	dog.Bark()
	dogPack.Put(dog)
}
