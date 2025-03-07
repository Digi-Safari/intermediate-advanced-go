package middleware

import (
	"errors"
	"rest-api/auth"
)

type Mid struct {
	a *auth.Auth
}

func NewMid(a *auth.Auth) (*Mid, error) {
	if a == nil {
		return nil, errors.New("auth cannot be nil")
	}
	return &Mid{a: a}, nil
}
