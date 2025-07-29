package repository

import (
	"context"

	"example.com/internal-service/internal/domain/user"
)

// UserRepository define as operações de persistência para User
type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	GetByID(ctx context.Context, id string) (*user.User, error)
	GetAll(ctx context.Context, page, limit int) ([]*user.User, int, error)
	Update(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id string) error
}
