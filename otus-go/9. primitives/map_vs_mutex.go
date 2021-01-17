package main

import "sync"

func main() {
	var wg sync.WaitGroup

	m := make(map[string]int)

	/**
	fatal error: concurrent map writes
	т.е. одновременная запись map!! такая ошибка вас ждет)
	ибо писать одновременно в map нельзя!!!
	*/
	for x := 0; x < 12; x++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			m["hello"] = 1
		}(&wg)
	}

	wg.Wait()
}
