package models

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CacheStore map[string]User

type Conn struct {
	store CacheStore
	//db *sql.DB
}

func (c *Conn) CreateUser(n NewUser) (User, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	user := c.store[n.Email]

	if user.Email == n.Email {
		return User{}, errors.New("user already exists")
	}

	us := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(passHash),
	}

	c.store = CacheStore{us.Email: us}
	return us, nil
}
