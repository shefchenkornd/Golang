package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//латинские одна буква = 1 байт
	//кириллические одна буква = 2 байта

	latin := "hello"      // здесь будет 5 байт
	cyrillic := "Привет!" // а здесь будет 13 байт, потому что на каждый кириллич.букву по 2 байта
	// объясню:
	// 6 кирил.букв (Привет) * 2 байта = 12 байт
	// один восклиц.знак * 1 байта = 1 байт
	// итого: 13 байт

	fmt.Printf("latin: %d byte, \nCyrillic: %d byte\n", len(latin), len(cyrillic))

	// если нужно включить в строку кавычки или переносы строки - используем обратные кавычки
	w := `hello
	"cruel"
	'world'
	`
	fmt.Println(w)

	//	## Что можно сделать со строками?

	// создавать
	s := "hello world \u9333 "

	// получить доступ к байту(!) в строке
	var c byte = s[0]
	fmt.Println("Получили байт из строки", c)

	// получить подстроку (в байтах!)
	var s2 string = s[0:5]
	fmt.Println("подстрока s[0:5] =", s2)

	// склеивать (это очень дорогая операция, старайтесь так не делать)
	s3 := s + "again"
	fmt.Println(s3)

	// узнавать длину в байтах
	length := len(s)
	fmt.Println("Длина строки:", length)

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	fmt.Println("------------------------------------------------------------")
	// ------------------------------------------------------------
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	fmt.Println("------------------------------------------------------------")

	const nihongo = "日本語"
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }

	fmt.Println("------------------------------------------------------------")
    
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
