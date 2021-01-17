package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DNAStrand2("ATTGC"))
}

var dnaReplacer *strings.Replacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

func DNAStrand2(dna string) string {
	return dnaReplacer.Replace(dna)
}