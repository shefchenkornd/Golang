package main

import (
	"fmt"
	"reflect"
)

type a struct{}
type b struct{}

/**
How to compare if two structs?
 */
func main() {
	aa := a{}
	bb := b{}

	fmt.Println(reflect.DeepEqual(aa, bb))
}
