package main

import "fmt"

func t1() error {
	for {}
}

func t2() error {
	for {}
}

func main() {
	chErr := make(ch error, 2) // буфер.канал с ёмкость два, чтобы хватило место для записи данных от двух функция

	go func() {
		chErr <- t1()
	}()

	go func() {
		chErr <- t2()
	}()

	// читаем из канала
	for err:= range chErr {
		fmt.Println(err)
	}

	// красиво завершить t1, t2 ....
}
