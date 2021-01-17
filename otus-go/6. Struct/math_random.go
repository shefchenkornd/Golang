package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("unix time =", time.Now().Unix())

	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("random []byte =", b)
}
