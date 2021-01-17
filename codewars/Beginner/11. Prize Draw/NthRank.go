package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

// To participate in a prize draw each one gives his/her firstname.
//
//Each letter of a firstname has a value which is its rank in the English alphabet. A and a have rank 1, B and b rank 2 and so on.
//
//The length of the firstname is added to the sum of these ranks hence a number som.
//
//An array of random weights is linked to the firstnames and each som is multiplied by its corresponding weight to get what they call a winning number.

func main() {
	st := "Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden"
	we := []int{1, 3, 5, 5, 3, 6}
	n := 2
	fmt.Println(NthRank(st, we, n))
}

const abc = "abcdefghijklmnopqrstuvwxyz"

var mapAlphabet map[string]int

type WinnerNumbers struct {
	m map[int][]string
}

func NthRank(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}

	namesParticipants := strings.Split(st, ",")
	if n > len(namesParticipants){
		return "Not enough participants"
	}

	mapAlphabet = createMapAlphabet()
	WinnerNumbers := newWinnerNumbers()

	for i, name := range namesParticipants {
		winnerNumber := getWinnerNumber(name, mapAlphabet, we[i])
		WinnerNumbers.add(winnerNumber, name)
	}

	keys := sortKeysDescOrder(*WinnerNumbers)

	counter := 1
	for _, key := range keys {
		sl := WinnerNumbers.m[key]
		for _, v := range sl {
			if n == counter {
				return v
			}
			counter++
		}
	}
	return ""
}

func sortKeysDescOrder(w WinnerNumbers) []int {
	var keys []int
	for key, _ := range w.m {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	return keys
}

func createMapAlphabet() map[string]int {
	m := make(map[string]int)
	for i, val := range abc {
		m[string(val)] = i+1
	}
	return m
}

func getRankByName(name string, m map[string]int) int {
	rank := 0

	name = strings.ToLower(name)
	for _, rune := range name {
		if el, ok := m[string(rune)]; ok {
			rank += el
		}
	}
	return rank
}

func getSom(name string, m map[string]int) int {
	return utf8.RuneCountInString(name) + getRankByName(name, m)
}

func getWinnerNumber(name string, m map[string]int, weigth int) int {
	return getSom(name, m) * weigth
}

// создание нового экземляра words
func newWinnerNumbers() *WinnerNumbers {
	return &WinnerNumbers{m: map[int][]string{}}
}

func (w *WinnerNumbers) add(number int, name string)  {
	sl, ok := w.m[number]

	if !ok {
		w.m[number] = []string{name}
		return
	}
	sl = append(sl, name)

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	w.m[number] = sl
}