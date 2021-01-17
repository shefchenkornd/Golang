package main

import "fmt"

// делаем бесконечный цикл for {}
//	for {
//		...
//		break
//		...
//	}
func main() {
	i := 0
	for {
		if i >= 5 {
			break // ВАЖНО!!! чтобы прервать бесконечный цикл используем ключевое слово `break`
		}

		fmt.Println(i)
		i++
	}
}
