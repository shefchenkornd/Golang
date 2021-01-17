package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"

	//"github.com/pkg/errors"
)

type error interface {
	Error() string
}

func Foor() error {
	return fmt.Errorf("My error code in %v", 32)
}

func main() {
	errors.New("IM an error")
	file, err := os.Open("file.go")

	defer func() {
		err := file.Close()
		if err != nil {
			errors.Wrap(err, "File close error")
		}
	}()

	err = Foor()
	if err != nil {
		fmt.Println(err)
	}
}
