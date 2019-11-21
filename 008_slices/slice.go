package main

import "fmt"

func main() {
	// slice или срезы динамически изменяет свой размер (вместимость)
	// len() длина слайса
	// cap() capacity / вместимость
	var sliceOne []int

	// создаём слайс и указывает capacity
	slice := make([]int, 5)

	fmt.Println(sliceOne)
	fmt.Println("slice", slice)
	fmt.Println("slice len", len(slice))
	fmt.Println("slice cap", cap(slice))

	sliceOne = append(sliceOne, 10)
	sliceOne = append(sliceOne, 20)
	sliceOne = append(sliceOne, 30)
	sliceOne = append(sliceOne, 40)
	sliceOne = append(sliceOne, 50)

	sliceOne[4] = 100
	// sliceOne[5] = 200 // panic: runtime error: index out of range [5] with length 5
	fmt.Println(sliceOne)

	// также можно нарезать слайсы
	fmt.Println(sliceOne[3:])
	fmt.Println(sliceOne[:4])
}
