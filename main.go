package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/marwell-dc/golang-chi-crud-api/routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok!"))
	})

	r.Mount("/books", routes.BookRoutes())
	http.ListenAndServe(":3000", r)
}
