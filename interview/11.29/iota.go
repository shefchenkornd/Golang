package main

import "fmt"

const (
	_ = iota
	Avito
	OLX
	LetGo
	GraigList = iota
	eBay
)

func main() {
	fmt.Println(Avito, OLX, LetGo, GraigList, eBay)
}
