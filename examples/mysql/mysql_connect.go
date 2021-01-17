package main

import (
	"database/sql"
	"time"
)

type Item struct{
	ID int
	Title string
}

/**
Пул соединений с базой данных это набор заранее открытых соединений с базой данных
 используемый для предоставления соединения в тот момент, когда оно требуется.
Пулы соединений используются для повышения производительности при работе с базами данных.
*/
func main() {
	db, err := sql.Open("mysql", "user:password@dbname")
	if err != nil {
		panic(err)
	}

	/*
		********************************
		See "Important settings" section
		*********************************
	*/

	/**
	Срок жизни соединения в пуле с момента его создания.
	Также проверка на возраст осуществляется каждый раз, когда мы получаем коннект из пула, и он тоже закывается
	, если истек срок его жизни
	 */
	db.SetConnMaxLifetime(time.Minute * 3)

	/**
	задает максимально возможное количество открытых соединений
	по умол = 0 - безлимитное кол-во соединений
	Обычно, после инициализации вызывают пинг, который создаёт коннект к БД
	err = db.Ping() -  подключение создается по требованию
	*/
	db.SetMaxOpenConns(10)

	/**
	управляем максимальным кол-вом соединений, которые могут храниться в пуле (DB.freeConn)
	Значение по умл. = 2
	После использования соединения оно освобождается и происходит попытка положить его в пул свободных коннектов.
	Если пул уже забит, то  освободившееся соединение закрывается и не попадает в пул.
	*/
	db.SetMaxIdleConns(10)

	authorId := 1
	item := []Item{}
	rows, err := db.QueryContext(ctx, "SELECT * FROM table where author_id=?", authorId)
	if err != nil {
		return nil, err
	}


	// ОБЯЗАТЕЛЬНО!!!
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		item := &Item{}
		err = rows.Scan(&item.ID, &item.Title)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

}
