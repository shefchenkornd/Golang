package main

import "fmt"

func main() {
	var runeCount, byteCount int

	str := "Hello 世界"

	for _ = range str {
		runeCount++
	}

	for _ = range []byte(str) {
		byteCount++
	}

	fmt.Println(runeCount, byteCount) // 8, 12
}
