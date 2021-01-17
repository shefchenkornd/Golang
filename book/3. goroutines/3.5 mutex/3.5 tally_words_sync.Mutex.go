package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// счётчик слов в состоянии гонки
func main() {
	var wg sync.WaitGroup

	w := newWords()

	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("\nWords that appear more than once:")
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

// структура облегчит реорганизацию кода
type words struct {
	found map[string]int
	sync.Mutex
}

// создание нового экземляра words
func newWords() *words {
	return &words{found: map[string]int{}}
}

// открытие файла, анализ содержимого и подсчёт найденных в нём слов.
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// сканнер - полезный инструмент для подобного анализа
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		//dict.mux.Lock()
		dict.add(word, 1)
		//dict.mux.Unlock()
	}
	return scanner.Err()
}

// фиксирует кол-во вхождений этого слова. Если слова еще нет, то добавляем его
func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}