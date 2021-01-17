package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// чтобы запустить через терминал
// go run 3.4\ compress_waitGroup.go exampledata/*
func main() {
	fmt.Println(time.Now().Format("2006-01-01 15:11:22.0000"))

	var wg sync.WaitGroup // нет необходимости инициализировать WaitGroup
	const AddDeltaForWaitGroupCounter = 1

	var i int = -1
	var file string

	for i, file = range os.Args[1:] {
		fmt.Println(i)
		wg.Add(AddDeltaForWaitGroupCounter) // для каждого файла сообщить группе, что ожидается выполнение еще одной операции сжатия

		// эта анонимная функция вызывает горутину и уведомляет группу ожидать её завершения
		// объявление горутины не приведет к её немедленному выполнению (из-за планировщика Го), поэтому
		// мы создали анонимную функцию и передаём ей уникальное значение переменной file, это является своего рода ГАРАНТИЕЙ
		// всякий раз запуская горутины в цикле и передавая ей переменную, создавайте копию этой переменной внутри цикла
		go func(filename string) {
			compress2(filename)
			wg.Done()
		}(file)
	}

	wg.Wait() // внешняя горутина (main) ожидает, пока все горутины, выполняющие сжатие, вызовут wg.Done()
	fmt.Printf("Compressed %d files\n", i+1)

	fmt.Println(time.Now().Format("2006-01-01 15:11:22.0000"))
}

func compress2(filename string) error {
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
