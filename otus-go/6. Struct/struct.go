package main

import "fmt"

var counter int

var wordCounts []struct {
	w string
	n int
}

var resp struct {
	Ok        bool `json: "ok"`
	Total     int  `json: "total"`
	Documents []struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}
}

type User struct {
	UserId   int64
	Login    string
	Password string
	previos  *User // рекурсивная структура ТОЛЬКО через pointer
	counter  *int
}

type Signal struct{}

func main() {
	s := struct{}{}
	user := User{} // Zero value для типа User
	user2 := &User{} // Тоже, но указатель
	user3 := User{
		1,
		"Vasya",
		"password",
		nil,
		nil,
	} // по номерам полей

	fmt.Printf("%T - %v\n", s, s)
	fmt.Printf("%T - %v\n", user, user)
	fmt.Printf("%T - %v\n", user2, user2)
	fmt.Printf("%T - %v\n", user3, user3)
}
