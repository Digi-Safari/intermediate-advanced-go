package middleware

import (
	"errors"
	"rest-api/auth"
)

// Whenever we need to inject a dependency for a package, we can create a struct,
// and then we can add the required struct as field, like what we have done for the mid struct

type Mid struct {
	a *auth.Auth
}

func NewMid(a *auth.Auth) (*Mid, error) {
	if a == nil {
		return nil, errors.New("auth cannot be nil")
	}
	return &Mid{a: a}, nil
}
