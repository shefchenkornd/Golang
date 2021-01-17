package main

import "fmt"

func main() {
	fmt.Println(DNAStrand("ATTGC"))
}

func DNAStrand(dna string) string {
	var convert string

	for _, char := range dna {
		if res, ok := complements[string(char)]; ok {
			convert += res
		}
	}

	return convert
}

type SymbolsComplements map[string]string

var complements = SymbolsComplements{
	"A": "T",
	"T": "A",
	"C": "G",
	"G": "C",
}
