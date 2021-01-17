package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"
)

// чтобы запустить через терминал
// go run 3.3\ compress.go exampledata/*
func main() {
	fmt.Println(time.Now().Format("2006-01-01 15:11:22.0000"))
	for _, file := range os.Args[1:] {
		compress(file)
	}

	fmt.Println(time.Now().Format("2006-01-01 15:11:22.0000"))
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}

	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
