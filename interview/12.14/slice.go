package main

import "fmt"

func main() {
	x := []int{0, 0}
	i := 0
	i, x[i] = 1, 2 // отработает вначале слайс, а потом переменная "i"
	/**
	i = 0
	x[i=0] = 2
	x = [2, 0]
	а затем произойдет присвоение i=1
	 */
	fmt.Println(x) // [2 0]
}