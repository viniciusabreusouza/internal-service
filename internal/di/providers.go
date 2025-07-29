package di

import (
	"context"

	"example.com/internal-service/internal/infra/grpc"
	http "example.com/internal-service/internal/infra/http"
	"example.com/internal-service/internal/repository"
	"example.com/internal-service/internal/service"
	"go.uber.org/zap"
)

// Providers para dependências básicas
func ProvideContext() context.Context {
	return context.Background()
}

func ProvideLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

// Providers para repositórios
func ProvideUserRepository() repository.UserRepository {
	return repository.NewUserMemoryRepository()
}

// Providers para serviços
func ProvideUserService(userRepo repository.UserRepository, log *zap.Logger) service.UserService {
	return service.NewUserService(userRepo, log)
}

// Providers para servidores
func ProvideHTTPServer(log *zap.Logger, userService service.UserService) (http.Server, error) {
	return http.NewHTTPServer(log, userService)
}

func ProvideGRPCServer(log *zap.Logger, userService service.UserService) (grpc.Server, error) {
	return grpc.NewGRPCServer(log, userService)
}
