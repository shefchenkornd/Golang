package main

/**
Интерфейс может встраивать другой интерфейс
 */
type Greeter interface {
	hello()
}

type Stranger interface {
	Greeter
	Bye() string
}

func main() {
	
}
