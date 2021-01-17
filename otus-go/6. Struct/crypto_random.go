package main

import (
	"crypto/rand"
	"fmt"
)

/**
Чек-лист безопасности для Go-разработчика. Елена Граховац, N26.
 */
func main() {
	b := make([]byte, 5)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Правильная КРИПТОГРАФИЧЕСКАЯ случайность", b)
}
