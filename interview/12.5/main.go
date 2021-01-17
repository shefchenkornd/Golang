package main

import "fmt"

func add(arr []int, v int)  {
	arr = append(arr, v)
}

func main() {
	var arr = make([]int, 0, 1000)
	fmt.Printf("%v %p\n", arr, &arr)
	add(arr, 10)
	fmt.Printf("%v %p\n", arr, &arr)
}
