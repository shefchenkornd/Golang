package main

import "fmt"

func main() {
	// rune - int 32 для Unicode
	// byte - int 8 для бинарных данных или англ.алфавита

	fmt.Printf("%c", 65) // является кодом заглавной буквы А
	fmt.Printf("%c", 1285515) // смайлик 😃

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang) // Выводит: πάω!

	// Запоминать все Юникоды нет нужды, в Go есть символьный литерал. Просто поместите символ в одинарные кавычки: 'A'.
	// Если тип не уточняется, Go присвоит тип rune. Следующие три строки кода эквиваленты:
	grade := 'A'
	fmt.Println(grade) // выведет код - 65

}
