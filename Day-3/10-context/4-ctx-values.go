package main

import (
	"context"
	"fmt"
)

// The provided key must be comparable and should not be of type string
// or any other built-in type to avoid collisions between packages using context.
// Users of WithValue should define their own types for keys.
type ctxKey string

const K ctxKey = "key"

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, K, "123")
	getRequestID(ctx)

}

func getRequestID(ctx context.Context) {
	// fetching the value from the context
	// using type assertion// making sure the value is of correct type
	//and interface is not nil
	reqId, ok := ctx.Value(K).(string)
	if !ok {
		fmt.Println("reqId not found or invalid type")
		return
	}
	fmt.Println(reqId)

}
