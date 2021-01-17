package main

import "fmt"

func main() {
	var p *int
	i := 948

	p = &i // point to i
	fmt.Println(p)

	// доступ к переменной i через указатель *p
	fmt.Println(*p) // прочитать i через указатель p

	*p = 21         // установить i через указатель p

	fmt.Println(i)
	fmt.Println(*p)
}
