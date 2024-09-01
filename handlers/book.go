package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marwell-dc/golang-chi-crud-api/models"
)

type BookHandler struct{}

var book *models.Book = &models.Book{}

func (b *BookHandler) NewBookHander() *BookHandler {
	return &BookHandler{}
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	var listBooks = book.ListBooks()
	err := json.NewEncoder(w).Encode(listBooks)

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := book.GetBookById(id)

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}

	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookCreate models.Book
	err := json.NewDecoder(r.Body).Decode(&bookCreate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.StoreBook(bookCreate)
	err = json.NewEncoder(w).Encode(bookCreate)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var bookUpdate models.Book
	err := json.NewDecoder(r.Body).Decode(&bookUpdate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook := book.Update(id, bookUpdate)

	if updatedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedBook := book.Delete(id)

	if deletedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
