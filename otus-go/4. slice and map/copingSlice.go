package main

import "fmt"

/**
при копировании слайса (а также получении под-слайса и передаче в функцию)
копируется только заголовок. Область памяти остается общей.
Но только до тех пор пока один из слайсов не "вырастет" (произведет реаллокацию)
*/
func main() {
	arr := []int{1, 2}

	arr2 := arr
	arr2[0] = 42
	fmt.Println(arr[0]) // 42, потому что сейчас оба слайса ссылаются на одну и туже область памяти

	arr2 = append(arr2, 3, 4, 5, 6)
	arr2[0] = 1
	fmt.Println(arr[0]) // после того как мы превысили ёмкость исходного слайса, то для второго слайса создался новый слайс с новым участком в памяти
	// и теперь эти два слайса имеют разные области в & &памяти.
	// см. картинку "неочевидные_следствия_после_копирования_слайса.png"

	/**
	если хотите написать фунцкию изменяющую слайс, сделайте так чтобы он возвращал новый слайс. Не изменяйте слайсы, которые передали вам как аргументы, т.к. это shalow (поверностная) копия исходных слайсов/
	Менять слайс, тот который вам передали, внутри функции БЕЗОПАСНО только если вы его НЕ УВЕЛИЧИВАЕТЕ, напр.: переставлять элементы.
	Если же вы будете увеличивать размер слайса, то с большой долей вероятности эти изменения потеряются, потому что они будут скопированы в новую область памяти, чтобы этого не произошло вам нужно возвращать новый результирующий слайс
	*/

	/**
	Если хотите получить полную копию слайса, используйте функцию `copy`
	*/
	s := []int{1, 2, 3}
	s2 := make([]int, len(s))
	copy(s2, s)
	s2[0] = 12

	fmt.Println(s2, s)
}