package main

import "fmt"

/**
Также с помощью `fmt` можно считывать данные в заранее известном формате
func Scanf(format string, a...interface{})
func Fscanf(r io.Reader, format string, a...interface{})
 */
func main() {
	var s string
	var d int64
	fmt.Scanf("%s %d", &s, &d) // нужно передавать POINTER (!!!) на пеерменные
}
