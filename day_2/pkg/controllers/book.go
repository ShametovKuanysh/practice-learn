package controllers

import (
	"day_2/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book := models.GetBookByID(uint(id))
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	b := book.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book := &models.Book{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	book.ID = uint(id)
	updatedBook := models.UpdateBook(book, uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	models.DeleteBook(uint(id))
	w.WriteHeader(http.StatusOK)
}
