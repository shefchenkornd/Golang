package main

import (
	"fmt"
	"reflect"
)

func main() {
	// используем пустой интерфейс, когда нам заранее неизвестно какой будет тип
	m := map[string]interface{}{}

	m["one"] = 1
	m["two"] = 2.0
	m["tree"] = true
	m["four"] = "school"

	for k, val := range m {
		switch val.(type) {
		case int:
			fmt.Printf("%s is integer\n", k)
		case float64:
			fmt.Printf("%s is float\n", k)
		default:
			fmt.Printf("%s is %v\n", k, reflect.TypeOf(val)) // применили рефлексию для определения типа переменной
		}
	}
}
