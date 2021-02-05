package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := []byte(`{"Name":"Vasya","Job":{"Department":"Operations","Title":"Boss"}}`)

	var p2 interface{}
	err := json.Unmarshal(j, &p2)
	if err != nil {
		return
	}
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("p2 %v - %T\n", p2) // p2 map[Job:map[Department:Operations Title:Boss] Name:Vasya]
	//fmt.Printf("name=%s\n", p2.)

	/**
	здесь используем type casting - для чего?
	для того что `p2` переменная имеет значение, но не имеет тип (%!T(MISSING))
	поэтому для приведения типа используют TYPE CASTING.
	 */
	person, ok := p2.(map[string]interface{})
	if ok {
		fmt.Printf("name=%s\n", person["Name"])	// name=Vasya
	}
}
