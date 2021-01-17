package main

import "fmt"

type structer struct {
	a []int
	// slice - это структура из 3х элементов
}

func (s *structer) Change()  {
	s.a = append(s.a, 100500)
	s.a[0] = 1111
}


func main() {
	a := structer{a: []int{1,2,3,4,5}}
	fmt.Println(a)

	a.Change()
	fmt.Println(a)
}
