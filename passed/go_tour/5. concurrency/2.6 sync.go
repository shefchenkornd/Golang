package main

import "fmt"

func main() {
	results := make(map[int]int)
	structCh := make(chan struct{})

	go factoriall(5, results, structCh) // Go-процедуры исполняются в том же адресном пространстве, поэтому доступ к общей памяти должен быть синхронизирован.

	<-structCh        // ожидаем закрытия канала structCh

	for i, v := range results {
		fmt.Println(i, " - ", v)
	}
}

func factoriall(n int, results map[int]int, structCh chan struct{}, ) {
	defer close(structCh)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
