package main

import (
	"fmt"
	"strconv"
)

func parseInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, nil
		}

		if b {
			n = 1
		}
	}

	return n, err
}

func main() {
	fmt.Println(parseInt("true"))
}
