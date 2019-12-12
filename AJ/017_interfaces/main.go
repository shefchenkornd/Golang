package main

import "fmt"

/**
ПОЛИМОРФИЗМ
- возможность объектов с одинаковой спецификацией иметь различную реализацию
*/

// описывает только методы, в них не должно быть полей/свойств
type PrintInterface interface {
	checkDate()
}

func (d *Document) checkDate()  {
	fmt.Println("checkDate for *Document")
}

func (pcard * PersonalCard) checkDate()  {
	fmt.Println("checkDate for *PersonalCard")
}

type Document struct {
	Date          string
	Number        string
	NumberOfPages int
}

type PersonalCard struct {
	Date      string
	FirstName string
	LastName  string
	Age       int
}

func main() {
	doc := new(Document)
	pcard := new(PersonalCard)

	doc.Date = "1.12.2019"
	doc.NumberOfPages = 5
	doc.Number = "A - 100"

	pcard.Date = "3.12.2019"
	pcard.Age = 27
	pcard.FirstName = "John"
	pcard.LastName = "Burg"

	sliceInterface := []PrintInterface{doc, pcard}

	PrintAnyDoc(sliceInterface)
}

func PrintAnyDoc(anyDocs []PrintInterface) {
	for _, v := range anyDocs {
		fmt.Println(v)
	}
}
