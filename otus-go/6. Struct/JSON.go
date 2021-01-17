package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Surname string `json:"surname,omitempty"`
	Fathername string `json:"fathername"`
	Age int `json:"-"`
}

func main() {
	person := Person{
		Name:       "Ilya",
		Surname:    "",
		Fathername: "Andrew",
		Age:        31,
	}

	json, _ := json.Marshal(person)
	fmt.Printf("%v", string(json))
}
