package main

import "fmt"

// ------------------
type Employee struct {
	name, surname string
}

// без указателю мы просто выводим данные содержащиеся в типе
func (e Employee) FullName() string {
	return e.name + " " + e.surname
}

// ------------------
type MyInt int

// по указателю мы меняем содержимое типа
func (i *MyInt) Print()  {
	*i++
	fmt.Println(*i)
}

// ------------------

func main() {
	fmt.Println(Employee{"Alexander", "Davidov"}.FullName())

	var i MyInt = 3

	i.Print()
	i.Print()
	i.Print()
}
