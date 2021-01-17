package main

import "fmt"

func main() {
	intCh := make(chan int, 3)

	intCh <- 1
	intCh <- 2

	close(intCh) // После закрытия канала мы не сможем послать в него новые данные


	fmt.Println(<-intCh) // Однако мы можем получить ранее добавленные данные.
	fmt.Println(<-intCh)
	fmt.Println(<-intCh) // Но при попытке получить данные из канала, которых нет, мы получим значение по умолчанию, те. 0
}
