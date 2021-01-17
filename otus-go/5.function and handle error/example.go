package main

import "fmt"

func studentName(name string) string {
	type Student struct {
		name string
	}

	s := Student{name}

	return s.name
}

func main() {
	println(studentName("Olga"))
	fmt.Println(111)
}
