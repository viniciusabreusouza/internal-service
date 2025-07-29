package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// User representa a entidade de domínio User
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser cria uma nova instância de User
func NewUser(name, email string) (*User, error) {
	if err := validateUserData(name, email); err != nil {
		return nil, err
	}

	now := time.Now()
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update atualiza os dados do usuário
func (u *User) Update(name, email string) error {
	if err := validateUserData(name, email); err != nil {
		return err
	}

	u.Name = name
	u.Email = email
	u.UpdatedAt = time.Now()
	return nil
}

// validateUserData valida os dados do usuário
func validateUserData(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	if len(name) > 100 {
		return errors.New("name cannot exceed 100 characters")
	}
	return nil
}
