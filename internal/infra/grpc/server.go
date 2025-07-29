package grpc

import (
	"context"
	"fmt"
	"net"

	"example.com/internal-service/internal/handler"
	"example.com/internal-service/internal/proto/user"
	"example.com/internal-service/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server interface {
	Run(context.Context) error
	Stop(context.Context) error
}

type GRPCServer struct {
	server *grpc.Server
	log    *zap.Logger
	port   int
}

func NewGRPCServer(log *zap.Logger, userService service.UserService) (*GRPCServer, error) {
	grpcServer := grpc.NewServer()

	// Registrar serviços
	userHandler := handler.NewUserGRPCHandler(userService, log)
	user.RegisterUserServiceServer(grpcServer, userHandler)

	server := GRPCServer{
		log:    log,
		server: grpcServer,
		port:   9090,
	}

	return &server, nil
}

func (s *GRPCServer) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.log.Info("Starting gRPC server", zap.Int("port", s.port))
	return s.server.Serve(lis)
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.log.Info("Stopping gRPC server")
	s.server.GracefulStop()
	return nil
}
