package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("Временный файл")

	// файл будет создан в os.TempDir, например, /tmp/exapmle-Jsm33
	tmpfile, err := ioutil.TempFile("", "example-")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // не забываем удалять
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
