package main

import (
	"fmt"
	"os"
)

// os.Args[0] всегда будет лежать Имя временного исполняемого файла, созданного командой go run
// os.Args[1] будут лежать аргументы переданные пользователем
// поэтому проверка мы делаем от "> 1"
func main() {
	fmt.Println("Name of the temporary executable created by the go run command: ", os.Args[0])

	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("hello, Gopher!")
	}
}
