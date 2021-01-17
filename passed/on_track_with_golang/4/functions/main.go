package main

import (
	"fmt"
	"time"
)

// напишем программу, которая будет возвращать строку в зависимости от текущего часа!
func main() {
	hourOfDay := time.Now().Hour() // int типа, 20 часов вечера
	greeting := getGreeting(hourOfDay)

	fmt.Println(greeting)
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning"
	} else if hour < 18 {
		return "Good Aftenoon"
	} else {
		return "Good Evening"
	}
}
