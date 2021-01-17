package main

import "fmt"

func main() {
	var arr [3]string

	arr[0] = "Go"
	arr[1] = "Java"
	arr[2] = "Ruby"
	//arr[3] = "LoL" // invalid array index 3 (out of bounds for 3-element array)

	fmt.Print(arr)
}
