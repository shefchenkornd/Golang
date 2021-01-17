package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("x = ", x, " | y = ", y, " | x+y = ", x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for v, _ := range s {
		sum += v
	}

	c <- sum
}
