package http

import (
	"context"
	"net/http"

	"example.com/internal-service/internal/handler"
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

func NewHTTPServer(log *zap.Logger) (Server, error) {
	router := gin.Default()

	router.Use(gin.Recovery())

	router.GET("/health", handler.HealthCheckHandler)

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
