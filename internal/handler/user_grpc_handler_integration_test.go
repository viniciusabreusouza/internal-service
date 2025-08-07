//go:build integration
// +build integration

package handler

import (
	"context"
	"testing"

	domainUser "example.com/internal-service/internal/domain/user"
	protoUser "example.com/internal-service/internal/proto/user"
	"example.com/internal-service/internal/repository"
	"example.com/internal-service/internal/service"
	"example.com/internal-service/tests/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// setupTestGRPCHandler cria um handler gRPC de teste com dependências reais
func setupTestGRPCHandler(t *testing.T) (*UserGRPCHandler, *mongo.Client) {
	// Iniciar container MongoDB compartilhado
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Configurar repositório
	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)

	// Configurar service
	logger, _ := zap.NewDevelopment()
	userService := service.NewUserService(userRepo, logger)

	// Configurar handler
	handler := NewUserGRPCHandler(userService, logger)

	return handler, client
}

func TestUserGRPCHandler_Integration_CreateUser(t *testing.T) {
	handler, _ := setupTestGRPCHandler(t)

	// Limpar banco antes do teste
	client, _ := integration.StartSharedMongoContainer(t)
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	ctx := context.Background()

	tests := []struct {
		name         string
		request      *protoUser.CreateUserRequest
		expectError  bool
		expectedCode codes.Code
	}{
		{
			name: "successful creation",
			request: &protoUser.CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			expectError:  false,
			expectedCode: codes.OK,
		},
		{
			name: "empty name",
			request: &protoUser.CreateUserRequest{
				Name:  "",
				Email: "john@example.com",
			},
			expectError:  true,
			expectedCode: codes.Internal,
		},
		{
			name: "empty email",
			request: &protoUser.CreateUserRequest{
				Name:  "John Doe",
				Email: "",
			},
			expectError:  true,
			expectedCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := handler.CreateUser(ctx, tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)

				// Verificar código de erro gRPC
				if st, ok := status.FromError(err); ok {
					assert.Equal(t, tt.expectedCode, st.Code())
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.User)

				// Verificar dados do usuário criado
				assert.NotEmpty(t, response.User.Id)
				assert.Equal(t, tt.request.Name, response.User.Name)
				assert.Equal(t, tt.request.Email, response.User.Email)
				assert.NotNil(t, response.User.CreatedAt)
				assert.NotNil(t, response.User.UpdatedAt)
			}
		})
	}
}

func TestUserGRPCHandler_Integration_GetUser(t *testing.T) {
	handler, _ := setupTestGRPCHandler(t)

	// Limpar banco antes do teste
	client, _ := integration.StartSharedMongoContainer(t)
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	ctx := context.Background()

	// Criar um usuário primeiro
	testUser, err := domainUser.NewUser("Test User", "test@example.com")
	require.NoError(t, err)

	// Inserir usuário diretamente no banco
	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)
	err = userRepo.Create(ctx, testUser)
	require.NoError(t, err)

	tests := []struct {
		name         string
		request      *protoUser.GetUserRequest
		expectError  bool
		expectedCode codes.Code
	}{
		{
			name: "successful retrieval",
			request: &protoUser.GetUserRequest{
				Id: testUser.ID,
			},
			expectError:  false,
			expectedCode: codes.OK,
		},
		{
			name: "user not found",
			request: &protoUser.GetUserRequest{
				Id: "non-existent-id",
			},
			expectError:  true,
			expectedCode: codes.NotFound,
		},
		{
			name: "empty user ID",
			request: &protoUser.GetUserRequest{
				Id: "",
			},
			expectError:  true,
			expectedCode: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := handler.GetUser(ctx, tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)

				// Verificar código de erro gRPC
				if st, ok := status.FromError(err); ok {
					assert.Equal(t, tt.expectedCode, st.Code())
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.User)

				// Verificar dados do usuário
				assert.Equal(t, testUser.ID, response.User.Id)
				assert.Equal(t, testUser.Name, response.User.Name)
				assert.Equal(t, testUser.Email, response.User.Email)
			}
		})
	}
}

func TestUserGRPCHandler_Integration_ListUsers(t *testing.T) {
	handler, _ := setupTestGRPCHandler(t)

	// Limpar banco antes do teste
	client, _ := integration.StartSharedMongoContainer(t)
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	ctx := context.Background()

	// Criar múltiplos usuários
	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)

	for i := 1; i <= 3; i++ {
		testUser, err := domainUser.NewUser("User "+string(rune(i+'0')), "user"+string(rune(i+'0'))+"@example.com")
		require.NoError(t, err)
		err = userRepo.Create(ctx, testUser)
		require.NoError(t, err)
	}

	tests := []struct {
		name          string
		request       *protoUser.ListUsersRequest
		expectedCount int32
	}{
		{
			name: "list all users",
			request: &protoUser.ListUsersRequest{
				Page:  1,
				Limit: 10,
			},
			expectedCount: 3,
		},
		{
			name: "list with pagination",
			request: &protoUser.ListUsersRequest{
				Page:  1,
				Limit: 2,
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := handler.ListUsers(ctx, tt.request)

			assert.NoError(t, err)
			assert.NotNil(t, response)
			assert.Len(t, response.Users, int(tt.expectedCount))
			assert.Equal(t, int32(3), response.Total) // Total sempre deve ser 3
			assert.Equal(t, tt.request.Page, response.Page)
			assert.Equal(t, tt.request.Limit, response.Limit)
		})
	}
}
