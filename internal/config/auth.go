package config

import (
	"os"
)

// AuthConfig contém as configurações do Auth0
type AuthConfig struct {
	Domain   string
	Audience string
}

// NewAuthConfig cria uma nova configuração do Auth0
func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		Domain:   getEnvOrDefault("AUTH0_DOMAIN", "dev-wg5phitnbt1bkw42.us.auth0.com"),
		Audience: getEnvOrDefault("AUTH0_AUDIENCE", "https://dev-wg5phitnbt1bkw42.us.auth0.com/api/v2/"),
	}
}

// getEnvOrDefault retorna o valor da variável de ambiente ou o valor padrão
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
