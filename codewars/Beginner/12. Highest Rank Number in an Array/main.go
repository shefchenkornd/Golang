package main

import (
	"fmt"
	"sort"
)

// Complete the method which returns the number which is most frequent in the given input array.
// If there is a tie for most frequent number, return the largest number among them.
//
//Note: no empty arrays will be given.
func main() {
	//ar := []int{10, 12, 8,8, 12, 12,8, 6, 4, 10, 10}
	ar := []int{55, 191, 146, 165, 155, 20, 93, 94, 110, 238, 173, 83, 131, 243, 253}
	fmt.Println(HighestRank(ar))
}

func HighestRank(nums []int) int {
	mapFrequencies := make(map[int]int)

	for _, v := range nums {
		if el, ok := mapFrequencies[v]; ok {
			mapFrequencies[v] = el + 1
		} else {
			mapFrequencies[v] = 1
		}
	}

	sk := rankByWordCount(mapFrequencies)

	if withSameFrequency(sk) {
		var m int
		for i, e := range sk {
			if i==0 || m < e.Key {
				m = e.Key
			}
		}
		return m
	}

	//fmt.Println(sk)
	return sk[0].Key
}

func rankByWordCount(intFrequencies map[int]int) PairList {
	pl := make(PairList, len(intFrequencies))
	i := 0
	for k, v := range intFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func withSameFrequency(p PairList) bool {
	same := p[0].Value
	for i, _ := range p {
		if p[i].Value != same {
			return false
		}
	}

	return true
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
