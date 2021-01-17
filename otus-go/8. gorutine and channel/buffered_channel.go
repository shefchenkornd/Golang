package main

import "fmt"

// читаем пачкой из канала
func main() {
	var bufChan = make(chan int, 3)

	for i := 0; i < 3; i++ {
		bufChan <- i
		fmt.Println("Записано!")
	}

	/**
	fmt.Println(<-bufChan) // 0
	fmt.Println(<-bufChan) // 1
	fmt.Println(<-bufChan) // 2
	*/
	close(bufChan) // если убрать эту строчку кода, то будет deadlock

	// читаем пачкой из канала
	for x := range bufChan {
		fmt.Println(x)
	}


}
