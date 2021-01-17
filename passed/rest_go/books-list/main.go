package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"books-list/models"
	"books-list/controllers"
	"books-list/driver"

	_ "database/sql"
	_ "github.com/lib/pq"
	_ "github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init()  {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	//router.HandleFunc("/books/{id}", getBook).Methods("GET")
	//router.HandleFunc("/books", addBook).Methods("POST")
	//router.HandleFunc("/books", updateBook).Methods("PUT")
	//router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}








