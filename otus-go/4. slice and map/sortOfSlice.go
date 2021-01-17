package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

/**
для нативных типов есть уже функции сортировки:
sort.Ints()с
sort.String()
 */
func main() {
	s := []int{3, 2, 1}
	sort.Ints(s)
	fmt.Println(s)

	str := []string{"moon", "cruel", "abc"}
	sort.Strings(str)
	fmt.Println(str)

	/**
	а для слайсов с кастомными типами данных, нужно использовать
	sort.Slice()
	 */
	custom := []User{
		{"vasya", 19},
		{"petya", 18},
	}
	sort.Slice(custom, func(i, j int) bool {
		return custom[i].Age < custom[j].Age
	})
	fmt.Println(custom)
}
