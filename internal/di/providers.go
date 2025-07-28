package di

import (
	"context"

	http "example.com/internal-service/internal/infra/http"
	"go.uber.org/zap"
)

// Este arquivo contém exemplos de como adicionar providers para serviços e repositórios
// quando você implementar a arquitetura completa.

// Exemplo de providers para serviços (quando você implementar)
// func ProvideUserService(userRepo repository.UserRepository) service.UserService {
//     return service.NewUserService(userRepo)
// }

// Exemplo de providers para repositórios (quando você implementar)
// func ProvideUserRepository(db *sql.DB) repository.UserRepository {
//     return repository.NewUserRepository(db)
// }

// Exemplo de provider para banco de dados (quando você implementar)
// func ProvideDatabase(config *Config) (*sql.DB, error) {
//     return sql.Open("postgres", config.DatabaseURL)
// }

// Exemplo de provider para configuração (quando você implementar)
// func ProvideConfig() (*Config, error) {
//     return LoadConfig()
// }

// Providers para dependências básicas
func ProvideContext() context.Context {
	return context.Background()
}

func ProvideLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func ProvideHTTPServer(log *zap.Logger) (http.Server, error) {
	return http.NewHTTPServer(log)
}
