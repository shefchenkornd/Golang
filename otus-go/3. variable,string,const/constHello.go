package main

const HelloConst = 3

var HelloVar = 5

func main() {
	// константы доступны только во время компиляции
	print(HelloVar, HelloConst)
}

// Скомпилируем и посмотрим символы
// go build -o constHello.out constHello.go
// go tool nm constHello.out | grep Hello