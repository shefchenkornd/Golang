package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	urls := []string{
		"http://google.com",
		"http://golang.org",
		"https://yandex.ru",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(resp.Status)
			resp.Body.Close()
		}(url)
	}
	wg.Wait()

}
