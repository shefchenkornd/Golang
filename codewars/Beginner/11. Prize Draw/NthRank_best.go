package main

import (
	"fmt"
	"sort"
"strings"
)

func main() {
	st := "Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden"
	we := []int{1, 3, 5, 5, 3, 6}
	n := 2
	fmt.Println(NthRank2(st, we, n))
}

type participantes struct {
	nombre string
	numero int
}

func NthRank2(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}

	nombres := strings.Split(st, ",")

	if n > len(nombres) {
		return "Not enough participants"
	}

	var par []participantes

	for i, e := range nombres{
		par = append(par, participantes{e, getWining(strings.ToLower(e), we[i])})
	}

	sort.SliceStable(par, func(i, j int) bool {
		if par[j].numero == par[i].numero {
			return strings.ToLower(par[j].nombre) > strings.ToLower(par[i].nombre)
		}
		return par[j].numero < par[i].numero
	})

	return par[n-1].nombre
}

func getWining(str string, weight int) int {
	sArray := strings.Split(str, "")
	n := 0

	for _, e := range sArray{
		fmt.Println(e, e[0], e[0]-96)
		n += int(e[0])-96
	}

	return (n + len(sArray)) * weight
}