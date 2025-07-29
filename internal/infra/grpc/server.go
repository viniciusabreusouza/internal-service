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

// ServiceHandlers define todos os handlers de serviços
type ServiceHandlers struct {
	UserService service.UserService
	// BlogService service.BlogService  // Descomente quando implementar
	// PostService service.PostService  // Descomente quando implementar
}

// NewGRPCServer cria uma nova instância do servidor gRPC
func NewGRPCServer(log *zap.Logger, handlers ServiceHandlers) (*GRPCServer, error) {
	grpcServer := grpc.NewServer()

	// Registrar todos os serviços
	registerServices(grpcServer, handlers, log)

	server := GRPCServer{
		log:    log,
		server: grpcServer,
		port:   50051,
	}

	return &server, nil
}

// registerServices registra todos os serviços no servidor gRPC
func registerServices(grpcServer *grpc.Server, handlers ServiceHandlers, log *zap.Logger) {
	// Registrar UserService
	if handlers.UserService != nil {
		userHandler := handler.NewUserGRPCHandler(handlers.UserService, log)
		user.RegisterUserServiceServer(grpcServer, userHandler)
		log.Info("UserService registered")
	}

	// Registrar BlogService (quando implementar)
	// if handlers.BlogService != nil {
	//     blogHandler := handler.NewBlogGRPCHandler(handlers.BlogService, log)
	//     blog.RegisterBlogServiceServer(grpcServer, blogHandler)
	//     log.Info("BlogService registered")
	// }

	// Registrar PostService (quando implementar)
	// if handlers.PostService != nil {
	//     postHandler := handler.NewPostGRPCHandler(handlers.PostService, log)
	//     post.RegisterPostServiceServer(grpcServer, postHandler)
	//     log.Info("PostService registered")
	// }
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
