package handler

import (
	"context"

	"example.com/internal-service/internal/proto/user"
	"example.com/internal-service/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserGRPCHandler implementa o serviço gRPC UserService
type UserGRPCHandler struct {
	user.UnimplementedUserServiceServer
	userService service.UserService
	logger      *zap.Logger
}

// NewUserGRPCHandler cria uma nova instância do handler gRPC
func NewUserGRPCHandler(userService service.UserService, logger *zap.Logger) *UserGRPCHandler {
	return &UserGRPCHandler{
		userService: userService,
		logger:      logger,
	}
}

// CreateUser implementa o método CreateUser do UserService
func (h *UserGRPCHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	h.logger.Info("gRPC CreateUser called", zap.String("name", req.Name), zap.String("email", req.Email))

	domainUser, err := h.userService.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		h.logger.Error("Failed to create user via gRPC", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	protoUser := &user.User{
		Id:        domainUser.ID,
		Name:      domainUser.Name,
		Email:     domainUser.Email,
		CreatedAt: timestamppb.New(domainUser.CreatedAt),
		UpdatedAt: timestamppb.New(domainUser.UpdatedAt),
	}

	return &user.CreateUserResponse{User: protoUser}, nil
}

// GetUser implementa o método GetUser do UserService
func (h *UserGRPCHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	h.logger.Info("gRPC GetUser called", zap.String("id", req.Id))

	domainUser, err := h.userService.GetUser(ctx, req.Id)
	if err != nil {
		h.logger.Error("Failed to get user via gRPC", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	protoUser := &user.User{
		Id:        domainUser.ID,
		Name:      domainUser.Name,
		Email:     domainUser.Email,
		CreatedAt: timestamppb.New(domainUser.CreatedAt),
		UpdatedAt: timestamppb.New(domainUser.UpdatedAt),
	}

	return &user.GetUserResponse{User: protoUser}, nil
}

// ListUsers implementa o método ListUsers do UserService
func (h *UserGRPCHandler) ListUsers(ctx context.Context, req *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	h.logger.Info("gRPC ListUsers called", zap.Int32("page", req.Page), zap.Int32("limit", req.Limit))

	page := int(req.Page)
	limit := int(req.Limit)

	domainUsers, total, err := h.userService.ListUsers(ctx, page, limit)
	if err != nil {
		h.logger.Error("Failed to list users via gRPC", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	protoUsers := make([]*user.User, len(domainUsers))
	for i, domainUser := range domainUsers {
		protoUsers[i] = &user.User{
			Id:        domainUser.ID,
			Name:      domainUser.Name,
			Email:     domainUser.Email,
			CreatedAt: timestamppb.New(domainUser.CreatedAt),
			UpdatedAt: timestamppb.New(domainUser.UpdatedAt),
		}
	}

	return &user.ListUsersResponse{
		Users: protoUsers,
		Total: int32(total),
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

// UpdateUser implementa o método UpdateUser do UserService
func (h *UserGRPCHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	h.logger.Info("gRPC UpdateUser called", zap.String("id", req.Id))

	domainUser, err := h.userService.UpdateUser(ctx, req.Id, req.Name, req.Email)
	if err != nil {
		h.logger.Error("Failed to update user via gRPC", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	protoUser := &user.User{
		Id:        domainUser.ID,
		Name:      domainUser.Name,
		Email:     domainUser.Email,
		CreatedAt: timestamppb.New(domainUser.CreatedAt),
		UpdatedAt: timestamppb.New(domainUser.UpdatedAt),
	}

	return &user.UpdateUserResponse{User: protoUser}, nil
}

// DeleteUser implementa o método DeleteUser do UserService
func (h *UserGRPCHandler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	h.logger.Info("gRPC DeleteUser called", zap.String("id", req.Id))

	err := h.userService.DeleteUser(ctx, req.Id)
	if err != nil {
		h.logger.Error("Failed to delete user via gRPC", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	return &user.DeleteUserResponse{Success: true}, nil
}
