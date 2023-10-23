package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

func New(email string, active bool) User {
	id := uuid.New()

	return User{
		Id:     id.String(),
		Email:  email,
		Active: active,
	}
}
