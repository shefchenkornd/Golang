package main

import "errors"

func main() {
	// также можем передать ошибку
	panic(errors.New("ooops"))

	// мы можем передать, как пустой интерфейс
	panic(nil)
}
