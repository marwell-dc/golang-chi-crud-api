package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marwell-dc/golang-chi-crud-api/handlers"
)

func BookRoutes() chi.Router {
	r := chi.NewRouter()

	bookHandler := handlers.BookHandler{}

	r.Get("/", bookHandler.ListBooks)
	r.Get("/{id}", bookHandler.GetBookById)
	r.Post("/", bookHandler.CreateBook)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)

	return r
}
