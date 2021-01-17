package main

import "fmt"

/**
nill != nill

вы спросите почему так?
ответ очень просто, скорее всего вы пытаетесь сравнить какой-то err == nill
если так, то вы должны знать, что:
тип Error
имеет Type and Value
 */
func handle() error {
	var err *echo.HTTPError = nill
	return err  // --------------> Type => *echo.HTTPError |  Value => nill
}

func main() {
	err := handle()

	// НЕПРАВИЛЬНО
	fmt.Println(err == nil) // false

	// ПРАВИЛЬНО
	fmt.Println(err == (*echo.HTTPError)nill) // true
}
