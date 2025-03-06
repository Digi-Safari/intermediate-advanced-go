package models

import (
	"errors"
)

var users = map[uint64]User{
	123: {
		FName:    "Bob",
		LName:    "abc",
		Password: "123456",
		Email:    "bob@email.com",
	},
}

var ErrUserNotFound = errors.New("user not found with the id ")

// GetUser fetches the data from the map
func GetUser(id uint64) (User, error) {
	u, ok := users[id]
	if !ok {

		return User{}, ErrUserNotFound
	}

	return u, nil
}
