package main

import "fmt"

func main() {
	intCh := make(chan int, 3)

	intCh <- 3
	intCh <- 5
	intCh <- 7
	// intCh <- 15  // блокировка - функция main ждет, когда освободится место в канале

	fmt.Println("Ёмкость канала =", cap(intCh))
	fmt.Println("количество элементов в канале =", len(intCh))

	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)



}
