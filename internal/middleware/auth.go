package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/internal-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	log    *zap.Logger
	config *config.AuthConfig
}

func NewAuthMiddleware(log *zap.Logger, authConfig *config.AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{
		log:    log,
		config: authConfig,
	}
}

// JWTAuthMiddleware valida tokens JWT do Auth0
func (am *AuthMiddleware) JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrair o token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			am.log.Warn("Authorization header missing")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Verificar se o header começa com "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			am.log.Warn("Invalid authorization header format")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header must start with 'Bearer '",
			})
			c.Abort()
			return
		}

		// Extrair o token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validar o token
		token, err := am.validateToken(tokenString)
		if err != nil {
			am.log.Warn("Token validation failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Invalid token: %v", err),
			})
			c.Abort()
			return
		}

		// Adicionar as claims do token ao contexto
		c.Set("token", token)
		c.Set("user_id", token.Subject())
		c.Set("audience", token.Audience())
		c.Set("issuer", token.Issuer())

		am.log.Info("Token validated successfully",
			zap.String("user_id", token.Subject()),
			zap.Strings("audience", token.Audience()),
		)

		c.Next()
	}
}

// validateToken valida o token JWT usando as chaves públicas do Auth0
func (am *AuthMiddleware) validateToken(tokenString string) (jwt.Token, error) {
	// URL do JWKS do Auth0
	jwksURL := fmt.Sprintf("https://%s/.well-known/jwks.json", am.config.Domain)

	// Carregar as chaves públicas do Auth0
	keySet, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}

	// Validar o token
	token, err := jwt.ParseString(tokenString,
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
		jwt.WithIssuer(fmt.Sprintf("https://%s/", am.config.Domain)),
		jwt.WithAudience(am.config.Audience),
		jwt.WithAcceptableSkew(30*time.Second), // Tolerância de 30 segundos para diferenças de relógio
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse/validate token: %w", err)
	}

	return token, nil
}
