package grpc

import (
	"context"
	"net"

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
}

func NewGRPCServer(log *zap.Logger) (*GRPCServer, error) {
	grpcServer := grpc.NewServer()

	server := GRPCServer{
		log:    log,
		server: grpcServer,
	}

	return &server, nil
}

func (s GRPCServer) Run(ctx context.Context) error {
	s.log.Info("Starting GRPC server")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s GRPCServer) Stop(ctx context.Context) error {
	s.log.Info("Stopping GRPC server")
	s.server.GracefulStop()
	return nil
}
