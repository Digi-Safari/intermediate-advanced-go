package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	mux := chi.NewRouter()
	mux.Use(middleware.Logger, middleware.Recoverer, mid)
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	// allowing specific kind of request only
	mux.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Hello, JSON! (Chi)",
			"status":  "success",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	//localhost:8080/user/123
	// Route Parameters
	mux.Get("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte("Hello, " + name + "! (Chi)"))
	})

	// localhost:8080/welcome?first_name="xyz"
	// Query Parameters
	mux.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		firstName := r.URL.Query().Get("first_name")
		if firstName == "" {
			firstName = "Guest"
		}
		lastName := r.URL.Query().Get("lastName")
		w.Write([]byte("Hello, " + firstName + " " + lastName + "!"))
	})

	http.ListenAndServe(":8080", mux)
}

func mid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
