package main

import "fmt"

// аналог класса
// в struct описываются только поля
type EDocument struct {
	Number string
	Date string
	NumberOfPages int
	Footer
}


type Footer struct {
	Author string
}

// создаем метод showAuthor() для EDocument
func (doc EDocument) showAuthor()  {
	fmt.Println(doc.Footer.Author)
}

// создаем метод showAuthor() для EDocument по ссылке на эту struct
func (doc *EDocument) showNumberOfPages()  {
	fmt.Println(doc.Footer.Author)
}

func main() {
	// инициализация класса - создаём объект
	doc1 := EDocument{
		Number:        "001-A",
		Date:          "10.12.2019",
		NumberOfPages: 10,
		Footer: Footer{Author: "Author-1"},
	}

	var doc2 EDocument
	doc2.Number = "002-A"
	doc2.Date = "10.12.2019"
	doc2.NumberOfPages = 22
	doc2.Author = "Author-2"

	// можно создавать объекты через ключевое слово new()
	// только объект передается по ссылке
	doc3 := new(EDocument)
	doc3.Number = "003-A"
	doc3.Date = "12.12.2019"
	doc3.NumberOfPages = 55
	doc3.Author = "Author-3"

	fmt.Printf("%T - %v \n", doc1, doc1)
	fmt.Printf("%T - %v \n", doc2, doc2)
	fmt.Printf("%T - %v \n", doc3, doc3)

	doc1.showAuthor()
	doc2.showAuthor()
	doc3.showAuthor()
}
