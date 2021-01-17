package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/**
	смещение
	 */
	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof("werdfdfdffsdfsdfdsfsdfsd"))

	type st struct {
		a byte
		b bool
		c uint64
	}

	fmt.Println(unsafe.Sizeof(st{}))
}
