package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	mux := chi.NewRouter()

	mux.Route("/v1/users", func(r chi.Router) {
		// get user
		r.Use(middleware.Logger)
		r.Get("/get", Mid(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "getting users")
		}))

		//get user by id
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "getting user by id")
		})

		// create one user
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {})
	})

	http.ListenAndServe(":8080", mux)
}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		next(w, r)
	}
}
