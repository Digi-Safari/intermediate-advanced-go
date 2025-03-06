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
		ctx := r.Context() // fetching the ctx object from the request

		id := uuid.NewString()
		// adding a new key value pair in the context
		ctx = context.WithValue(ctx, RequestIdKey, id)
		fmt.Println(id)
		//r.WithContext will update the existing context with the updated one
		next(w, r.WithContext(ctx)) // calling next thing in the chain
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	// fetching the ctx value using the key
	reqId, ok := ctx.Value(RequestIdKey).(string)
	if !ok {
		reqId = "unknown"
	}
	fmt.Println("hello handler", reqId)
	// in case if you
	AddToDb(ctx) // if other functions need context after handler function,
	// pass ctx as an argument
	fmt.Fprintln(w, "Hello")

}

func AddToDb(ctx context.Context) {
	//sql.DB{}.ExecContext(ctx)

}
