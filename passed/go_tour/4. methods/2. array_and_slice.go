package main

import "fmt"

type Item struct {
	i int
	b bool
}

func main() {
	arr := [3]int{2, 3, 4}

	// таким образом мы полностью передали в slice весь имеющийся массив
	var mySlice []int = arr[:]

	fmt.Println(arr)
	fmt.Println(mySlice, mySlice[:cap(mySlice)])
	fmt.Println(arr[:])



	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s, cap(s), len(s))

	item := Item{15, false}
	s = append(s, item)
	fmt.Println(s, cap(s), len(s))

	//a := make([]int, 5, 7)
	//fmt.Println(a, len(a), cap(a))
}
