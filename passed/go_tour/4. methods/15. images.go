package main

import (
	"fmt"
	"image"
	"runtime"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())


	runtime.GOMAXPROCS(1)

	cpu := runtime.NumCPU()
	fmt.Println(cpu)
}
