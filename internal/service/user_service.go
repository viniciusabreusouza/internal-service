package service

import (
	"context"

	"example.com/internal-service/internal/domain/user"
	"example.com/internal-service/internal/repository"
	"go.uber.org/zap"
)

// UserService define as operações de negócio para User
type UserService interface {
	CreateUser(ctx context.Context, name, email string) (*user.User, error)
	GetUser(ctx context.Context, id string) (*user.User, error)
	ListUsers(ctx context.Context, page, limit int) ([]*user.User, int, error)
	UpdateUser(ctx context.Context, id, name, email string) (*user.User, error)
	DeleteUser(ctx context.Context, id string) error
}

// userService implementa UserService
type userService struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

// NewUserService cria uma nova instância do UserService
func NewUserService(userRepo repository.UserRepository, logger *zap.Logger) UserService {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

// CreateUser cria um novo usuário
func (s *userService) CreateUser(ctx context.Context, name, email string) (*user.User, error) {
	s.logger.Info("Creating user", zap.String("name", name), zap.String("email", email))

	user, err := user.NewUser(name, email)
	if err != nil {
		s.logger.Error("Failed to create user domain object", zap.Error(err))
		return nil, err
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		s.logger.Error("Failed to save user to repository", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User created successfully", zap.String("id", user.ID))
	return user, nil
}

// GetUser busca um usuário pelo ID
func (s *userService) GetUser(ctx context.Context, id string) (*user.User, error) {
	s.logger.Info("Getting user", zap.String("id", id))

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get user from repository", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User retrieved successfully", zap.String("id", id))
	return user, nil
}

// ListUsers lista usuários com paginação
func (s *userService) ListUsers(ctx context.Context, page, limit int) ([]*user.User, int, error) {
	s.logger.Info("Listing users", zap.Int("page", page), zap.Int("limit", limit))

	// Validação de parâmetros
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, total, err := s.userRepo.GetAll(ctx, page, limit)
	if err != nil {
		s.logger.Error("Failed to get users from repository", zap.Error(err))
		return nil, 0, err
	}

	s.logger.Info("Users listed successfully", zap.Int("count", len(users)), zap.Int("total", total))
	return users, total, nil
}

// UpdateUser atualiza um usuário existente
func (s *userService) UpdateUser(ctx context.Context, id, name, email string) (*user.User, error) {
	s.logger.Info("Updating user", zap.String("id", id))

	// Buscar usuário existente
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get user for update", zap.Error(err))
		return nil, err
	}

	// Atualizar dados
	if err := existingUser.Update(name, email); err != nil {
		s.logger.Error("Failed to update user domain object", zap.Error(err))
		return nil, err
	}

	// Salvar no repositório
	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		s.logger.Error("Failed to save updated user to repository", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User updated successfully", zap.String("id", id))
	return existingUser, nil
}

// DeleteUser remove um usuário
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	s.logger.Info("Deleting user", zap.String("id", id))

	if err := s.userRepo.Delete(ctx, id); err != nil {
		s.logger.Error("Failed to delete user from repository", zap.Error(err))
		return err
	}

	s.logger.Info("User deleted successfully", zap.String("id", id))
	return nil
}
