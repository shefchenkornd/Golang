package main

import "fmt"

/**
data races - сложная ошибка для дебага
Условие, при котором порядок доступа к памяти определяет корректность результата!
 */
func main() {
	c := make(chan bool)
	m := make(map[string]string)

	go func() {
		m["1"] = "a" // first conflicting access
		c <- true
	}()

	<-c
	m["2"] = "b" // second conflicting access

	for k, v := range m {
		fmt.Println(k, v)
	}
}
