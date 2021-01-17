package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// The exit code 1 is a POSIX standard indicating the program has finished, but errors have occurred.
// напишем программу, которая будет возвращать строку в зависимости от текущего часа!
func main() {
	hourOfDay := time.Now().Hour() // int типа, 20 часов вечера
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // The exit code 1 is a POSIX standard indicating the program has finished, but errors have occurred.
	}

	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 7 {
		return message, errors.New("Too early for greeting")
	} else if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Aftenoon"
	} else {
		message = "Good Evening"
	}

	return message, nil
}
