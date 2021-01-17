package main

import "os"

func main() {
	path := "/tmp"
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не найден
		}

		// другие ошибки, например, нет прав
	}

	defer file.Close()
}
