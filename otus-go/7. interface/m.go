package main

import "fmt"

type Adult interface {
	IsAdult() bool
	fmt.Stringer
}

type Person struct {
	age int
	name string
}

func adultFilter(people []Adult) []Adult {
	adults := make([]Adult, 0)
	for _, p := range people {
		if p.IsAdult() {
			adults = append(adults, p)
		}
	}

	return adults
}

func (p Person) IsAdult() bool {
	if p.age < 18 {
		return false
	}
	return true
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %v, age: %v", p.name, p.age)
}

func main() {
	people := []Adult{
		Person{
			age:  15,
			name: "John",
		},
		Person{
			age:  18,
			name: "Joe",
		},
		Person{
			age:  45,
			name: "Mary",
		},
	}

	fmt.Println(adultFilter(people))
}
