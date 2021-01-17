package main

import (
	"fmt"
	"hello/model"
)


func main() {
	jumperList := model.GetList() // мы логику работы приложения разбили на пакеты и переместили в пакет model
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
