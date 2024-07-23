package model

import (
	"context"

	"github.com/google/uuid"
)

type GetUserOpts struct {
	UserID uuid.UUID
	Email  string
	Status UserStatus
}
type UserRepo interface {
	Get(opts GetUserOpts, ctx context.Context) (*User, error)
	Create(user User) error
}
