package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"books-list/models"
	"books-list/repository/book"
	"strconv"
)

type Controller struct {}

var books []models.Book

func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookRepository.BookRepository{}

		param := mux.Vars(r)
		id, err := strconv.Atoi(param["id"])
		logFatal(err)

		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookRepository.BookRepository{}
		var bookId int

		json.NewDecoder(r.Body).Decode(&book)
		bookId = bookRepo.AddBook(db, book, bookId)

		json.NewEncoder(w).Encode(bookId)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookRepository.BookRepository{}
		json.NewDecoder(r.Body).Decode(&book)

		rowsUpdated := bookRepo.UpdateBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		rowsDeleted := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
