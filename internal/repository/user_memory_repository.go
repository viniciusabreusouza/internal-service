package repository

import (
	"context"
	"errors"
	"sync"

	"example.com/internal-service/internal/domain/user"
)

// UserMemoryRepository implementa UserRepository usando memória
type UserMemoryRepository struct {
	users map[string]*user.User
	mutex sync.RWMutex
}

// NewUserMemoryRepository cria uma nova instância do repositório em memória
func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[string]*user.User),
	}
}

// Create salva um novo usuário
func (r *UserMemoryRepository) Create(ctx context.Context, user *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

// GetByID busca um usuário pelo ID
func (r *UserMemoryRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAll retorna todos os usuários com paginação
func (r *UserMemoryRepository) GetAll(ctx context.Context, page, limit int) ([]*user.User, int, error) {
	r.users = map[string]*user.User{
		"1": {
			ID:    "1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	total := len(r.users)
	if total == 0 {
		return []*user.User{}, 0, nil
	}

	// Calcular offset
	offset := (page - 1) * limit
	if offset >= total {
		return []*user.User{}, total, nil
	}

	// Converter map para slice
	users := make([]*user.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	// Aplicar paginação
	end := offset + limit
	if end > total {
		end = total
	}

	return users[offset:end], total, nil
}

// Update atualiza um usuário existente
func (r *UserMemoryRepository) Update(ctx context.Context, user *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

// Delete remove um usuário
func (r *UserMemoryRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
