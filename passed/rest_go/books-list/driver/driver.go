package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	/*
		pgUrl:
		- `dbname`
		- `host`
		- `password`
		- `port`
		- `user`
	*/
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}