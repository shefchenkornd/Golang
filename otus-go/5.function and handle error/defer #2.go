package main

import "fmt"

/**
к именованному аргументу мы имеем доступ поэтому defer изменит переменную `i`
так лучше не использовать на практике
 */
func c() (i int)  {
	defer func() { i++ }() // output: 2
	return 1
}

func main() {
	ll := c()
	fmt.Printf("%d", ll) // output: 2
}
