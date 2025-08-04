package http

import (
	"context"
	"net/http"

	"example.com/internal-service/internal/config"
	"example.com/internal-service/internal/handler"
	"example.com/internal-service/internal/middleware"
	"example.com/internal-service/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server interface {
	Run(context.Context) error
	Stop(context.Context) error
}

type HTTPServer struct {
	server *http.Server
	log    *zap.Logger
}

func NewHTTPServer(log *zap.Logger, userService service.UserService) (Server, error) {
	router := gin.Default()

	router.Use(gin.Recovery())

	// Health check
	router.GET("/health", handler.HealthCheckHandler)

	// User routes
	userHandler := handler.NewUserHTTPHandler(userService, log)
	users := router.Group("/api/v1/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/", userHandler.ListUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	// Protected routes
	authConfig := config.NewAuthConfig()
	authMiddleware := middleware.NewAuthMiddleware(log, authConfig)
	protectedHandler := handler.NewProtectedHandler(log)
	
	protected := router.Group("/")
	protected.Use(authMiddleware.JWTAuthMiddleware())
	{
		protected.GET("/dados-protegidos", protectedHandler.GetProtectedData)
	}

	log.Info("Routes configured", 
		zap.String("health", "/health"),
		zap.String("protected", "/dados-protegidos"),
	)

	server := HTTPServer{
		log: log,
	}

	s := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.server = &s

	return &server, nil
}

func (s HTTPServer) Run(ctx context.Context) error {
	s.log.Info("Starting HTTP server")
	s.log.Info("Server running on port 8080")

	return s.server.ListenAndServe()
}

func (s HTTPServer) Stop(ctx context.Context) error {
	s.log.Info("Stopping HTTP server")
	return s.server.Shutdown(ctx)
}
