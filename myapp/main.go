package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Books variable as a slice
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create a book
func creatBook(w http.ResponseWriter, r *http.Request) {

}

// Update the book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initialize mux router
	r := mux.NewRouter()

	// Mock data
	// TODO: implement database
	books = append(books, Book{ID: "1", ISBN: "122", Title: "The book one", Author: &Author{Firstname: "Jhon", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", ISBN: "53162", Title: "The book two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// endpoints setup
	r.HandleFunc("/api/books/", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", creatBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
