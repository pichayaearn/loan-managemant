package model

import "context"

type CreateUser struct {
	Email    string
	Password string
}

type UserService interface {
	GetUser(opts GetUserOpts, ctx context.Context) (*User, error)
	CreateUser(opts CreateUser) error
}
