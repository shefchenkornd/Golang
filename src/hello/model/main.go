package model

type jumper interface {
	Jump() string // метод интерфейса тоже должен быть экспортируемым, но НЕ интерфейсы и type struct
}

type gopher struct {
	name string
	age int
}

func (g gopher) Jump() string {
	return "gopher jump"
}

type horse struct {
	name string
	weight int
}

func (h horse) Jump() string {
	return "horse jump"
}

// для того чтобы функция была экспортируемая, необходимо написать название с Заглавной буквой
func GetList() []jumper {
	phil := &gopher{
		name: "phil",
		age:  25,
	}

	barbaro := &horse{
		name:   "barbaro",
		weight: 2000,
	}

	return []jumper{phil, barbaro}
}