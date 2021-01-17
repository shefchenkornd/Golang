package main

import (
	"fmt"
	"sort"
)

func main() {
	// сортировка maps через ключи (с помощью slice)!
	m := map[string]int{
		"0e1": 0,
		"1e1": 1,
		"2e1": 2,
		"3e1": 3,
		"4e1": 4,
		"5e1": 5,
	}

	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys{
		fmt.Println(k, " - ", m[k])
	}
}
