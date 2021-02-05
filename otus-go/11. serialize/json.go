package main

import (
	"encoding/json"
	"fmt"
)

// json теги помогают сделать алиасы к свойствам структуры
// omitempty = omit empty (опустить пустые)
// чтобы поле структуры не показывать в json указываем `json:"-"`
type Person struct {
	Name string `json:"fullname,omitempty"`
	Surname string `json:"familyname,omitempty"`
	Age  int	`json:"-"`
	Job  struct {
		Department string
		Title      string
	}
}

func main() {
	p1 := &Person{
		Name: "Name",
		Surname: "Surname",
		Age:  45,
		Job: struct {
			Department string
			Title      string
		}{Department: "Operations", Title: "Boss"},
	}

	j, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("p1 json %s\n", j) // p1 json {"Name":"Vasya","Job":{"Department":"Operations","Title":"Boss"}}

	var p2 Person
	json.Unmarshal(j, &p2)
	fmt.Printf("p2 %v\n", p2) // p2 {Vasya 0 {Operations Boss}}
}
