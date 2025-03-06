package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type contextKey string

const RequestIdKey contextKey = "request-id"

func main() {
	http.HandleFunc("/hello", RequestIdMid(Hello))
	http.ListenAndServe(":8080", nil)

}

func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := uuid.NewString()
		// adding a new key value pair in the context
		ctx = context.WithValue(ctx, RequestIdKey, id)
		fmt.Println(id)
		//r.WithContext will update the existing context with the updated one
		next(w, r.WithContext(ctx))
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
