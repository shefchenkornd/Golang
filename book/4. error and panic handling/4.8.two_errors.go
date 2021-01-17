package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrTimeout = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

var random = rand.New(rand.NewSource(time.Now().Unix()))

func sendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}

	return "", nil
}

func main() {
	response, err := sendRequest("Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

	fmt.Printf("%v", random.Int() %3)
}
