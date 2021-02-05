package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("start")
	cmd := exec.Command("sleep", "3")
	err := cmd.Run() // запускает команду и дожидается завершения
	if err != nil {
		return
	}
	fmt.Println("finish")

	err = cmd.Start() // запуск программы, без дожидания завершения
	if err != nil {
		return
	}
	cmd.Wait()         // дожидается завершения
}
