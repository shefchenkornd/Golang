package main

// интерфейс - это набор методов
// в данном примере - это интерфейс фигуры
type Shape interface {
	Area() float64
	Perimeter() float64
}

/**
утиная типизация - крякает как утка, и выглядит как утка, то это скорее всего утка
 */

/**
интерфейсы реализуются НЕЯВНО!!!!!
 */
type Duck interface {
	Talk() // разговаривает
	Swim() // плавает
	Walk() // ходит
}

/**
раз интерфейс Dog реализует методы Duck, то по мнению Го Dog - это будет Duck */
type Dog struct {
	name string
}

func (d Dog) Talk() string {
	return "AGGGGRRRRR"
}

func (d Dog) Walk()  {
	// do it
}

func (d Dog) Swim()  {
	// do it
}

func main() {
	
}
