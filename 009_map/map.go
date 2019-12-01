package main

import "fmt"

func main() {
	var simpleMap = map[int]string{}
	fmt.Println(simpleMap)

	simpleMap[198] = "some word"
	fmt.Println(simpleMap)

	var anotherMap = make(map[string]int)
	anotherMap["one"] = 1
	anotherMap["two"] = 2
	anotherMap["three"] = 3
	fmt.Println(anotherMap)

	delete(anotherMap, "two")
	fmt.Println(anotherMap)

	if value, ok := anotherMap["one"]; ok {
		fmt.Println(value)
	}
}
