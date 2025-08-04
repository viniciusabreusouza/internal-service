package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

type ProtectedHandler struct {
	log *zap.Logger
}

func NewProtectedHandler(log *zap.Logger) *ProtectedHandler {
	return &ProtectedHandler{
		log: log,
	}
}

// GetProtectedData retorna dados protegidos após validação do token
func (ph *ProtectedHandler) GetProtectedData(c *gin.Context) {
	// O token já foi validado pelo middleware
	token, exists := c.Get("token")
	if !exists {
		ph.log.Error("Token not found in context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token not found in context",
		})
		return
	}

	jwtToken, ok := token.(jwt.Token)
	if !ok {
		ph.log.Error("Invalid token type in context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid token type in context",
		})
		return
	}

	// Extrair informações do token
	userID, _ := c.Get("user_id")
	audience, _ := c.Get("audience")
	issuer, _ := c.Get("issuer")

	// Criar resposta com informações do token
	response := gin.H{
		"message": "Token válido! Acesso permitido.",
		"token_info": gin.H{
			"subject":    jwtToken.Subject(),
			"audience":   audience,
			"issuer":     issuer,
			"issued_at":  jwtToken.IssuedAt().Format(time.RFC3339),
			"expires_at": jwtToken.Expiration().Format(time.RFC3339),
		},
		"user_id": userID,
	}

	ph.log.Info("Protected data accessed successfully",
		zap.String("user_id", userID.(string)),
		zap.Strings("audience", audience.([]string)),
	)

	c.JSON(http.StatusOK, response)
}
