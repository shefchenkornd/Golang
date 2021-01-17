package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Teacher struct{}

func main() {
	var teacher Teacher

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	row := db.QueryRow("SELECT * FROM teachers WHERE id=?", 12)
	err = row.Scan(&teacher)
}
