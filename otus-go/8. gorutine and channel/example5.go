package main

import "fmt"

func main() {
	var start = make(chan struct{})

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			<-start
			// горутины не начнут работать, пока не будут созданы 100 горутин
		}(i)
	}

	close(start)
}
