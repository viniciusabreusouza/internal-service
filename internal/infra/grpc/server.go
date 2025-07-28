package grpc

import (
	"context"

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
