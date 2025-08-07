//go:build integration
// +build integration

package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/internal-service/internal/domain/user"
	"example.com/internal-service/internal/repository"
	"example.com/internal-service/internal/service"
	"example.com/internal-service/tests/integration"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// setupTestServer cria um servidor HTTP de teste com dependências reais
func setupTestServer(t *testing.T) (*gin.Engine, *mongo.Client) {
	// Configurar Gin para modo de teste
	gin.SetMode(gin.TestMode)

	// Iniciar container MongoDB
	_, client := integration.StartMongoContainer(t)

	// Configurar repositório
	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)

	// Configurar service
	logger, _ := zap.NewDevelopment()
	userService := service.NewUserService(userRepo, logger)

	// Configurar handler
	userHandler := NewUserHTTPHandler(userService, logger)

	// Configurar router
	router := gin.New()
	router.Use(gin.Recovery())

	// Configurar rotas
	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUser)
		api.GET("/users", userHandler.ListUsers)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return router, client
}

func TestUserHTTPHandler_Integration_CreateUser(t *testing.T) {
	router, client := setupTestServer(t)
	defer integration.CleanupMongoContainer(t, nil, client)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "successful creation",
			requestBody: map[string]interface{}{
				"name":  "John Doe",
				"email": "john@example.com",
			},
			expectedStatus: http.StatusCreated,
			expectedError:  false,
		},
		{
			name: "missing name",
			requestBody: map[string]interface{}{
				"email": "john@example.com",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "missing email",
			requestBody: map[string]interface{}{
				"name": "John Doe",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "invalid email format",
			requestBody: map[string]interface{}{
				"name":  "John Doe",
				"email": "invalid-email",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparar requisição
			jsonBody, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			// Executar requisição
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verificar status
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				// Verificar resposta de sucesso
				var response UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)

				assert.NotEmpty(t, response.ID)
				assert.Equal(t, tt.requestBody["name"], response.Name)
				assert.Equal(t, tt.requestBody["email"], response.Email)
				assert.NotEmpty(t, response.CreatedAt)
				assert.NotEmpty(t, response.UpdatedAt)
			} else {
				// Verificar resposta de erro
				var errorResponse map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
				require.NoError(t, err)

				assert.Contains(t, errorResponse, "error")
			}
		})
	}
}

func TestUserHTTPHandler_Integration_GetUser(t *testing.T) {
	router, client := setupTestServer(t)
	defer integration.CleanupMongoContainer(t, nil, client)

	// Criar um usuário primeiro
	user, err := user.NewUser("Test User", "test@example.com")
	require.NoError(t, err)

	// Inserir usuário diretamente no banco
	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)
	err = userRepo.Create(context.Background(), user)
	require.NoError(t, err)

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "successful retrieval",
			userID:         user.ID,
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "user not found",
			userID:         "non-existent-id",
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparar requisição
			url := fmt.Sprintf("/api/v1/users/%s", tt.userID)
			req := httptest.NewRequest("GET", url, nil)

			// Executar requisição
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verificar status
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				// Verificar resposta de sucesso
				var response UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)

				assert.Equal(t, user.ID, response.ID)
				assert.Equal(t, user.Name, response.Name)
				assert.Equal(t, user.Email, response.Email)
			} else {
				// Verificar resposta de erro
				var errorResponse map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
				require.NoError(t, err)

				assert.Contains(t, errorResponse, "error")
			}
		})
	}
}

func TestUserHTTPHandler_Integration_ListUsers(t *testing.T) {
	router, client := setupTestServer(t)
	defer integration.CleanupMongoContainer(t, nil, client)

	// Criar múltiplos usuários
	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)

	users := []*user.User{}
	for i := 1; i <= 3; i++ {
		testUser, err := user.NewUser(fmt.Sprintf("User %d", i), fmt.Sprintf("user%d@example.com", i))
		require.NoError(t, err)
		err = userRepo.Create(context.Background(), testUser)
		require.NoError(t, err)
		users = append(users, testUser)
	}

	tests := []struct {
		name           string
		query          string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "list all users",
			query:          "",
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "list with pagination",
			query:          "?page=1&limit=2",
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparar requisição
			url := fmt.Sprintf("/api/v1/users%s", tt.query)
			req := httptest.NewRequest("GET", url, nil)

			// Executar requisição
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verificar status
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Verificar resposta
			var response ListUsersResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Len(t, response.Users, tt.expectedCount)
			assert.Equal(t, 3, response.Total) // Total sempre deve ser 3
		})
	}
}

func TestUserHTTPHandler_Integration_UpdateUser(t *testing.T) {
	router, client := setupTestServer(t)
	defer integration.CleanupMongoContainer(t, nil, client)

	// Criar um usuário primeiro
	testUser, err := user.NewUser("Original User", "original@example.com")
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)
	err = userRepo.Create(context.Background(), testUser)
	require.NoError(t, err)

	tests := []struct {
		name           string
		userID         string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedError  bool
	}{
		{
			name:   "successful update",
			userID: testUser.ID,
			requestBody: map[string]interface{}{
				"name":  "Updated User",
				"email": "updated@example.com",
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:   "user not found",
			userID: "non-existent-id",
			requestBody: map[string]interface{}{
				"name":  "Updated User",
				"email": "updated@example.com",
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
		{
			name:   "invalid request body",
			userID: testUser.ID,
			requestBody: map[string]interface{}{
				"name": "", // Nome vazio
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparar requisição
			jsonBody, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			url := fmt.Sprintf("/api/v1/users/%s", tt.userID)
			req := httptest.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			// Executar requisição
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verificar status
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				// Verificar resposta de sucesso
				var response UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)

				assert.Equal(t, testUser.ID, response.ID)
				assert.Equal(t, tt.requestBody["name"], response.Name)
				assert.Equal(t, tt.requestBody["email"], response.Email)
			} else {
				// Verificar resposta de erro
				var errorResponse map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
				require.NoError(t, err)

				assert.Contains(t, errorResponse, "error")
			}
		})
	}
}

func TestUserHTTPHandler_Integration_DeleteUser(t *testing.T) {
	router, client := setupTestServer(t)
	defer integration.CleanupMongoContainer(t, nil, client)

	// Criar um usuário primeiro
	testUser, err := user.NewUser("Test User", "test@example.com")
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	userRepo := repository.NewUserMongoRepository(db)
	err = userRepo.Create(context.Background(), testUser)
	require.NoError(t, err)

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "successful deletion",
			userID:         testUser.ID,
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "user not found",
			userID:         "non-existent-id",
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Preparar requisição
			url := fmt.Sprintf("/api/v1/users/%s", tt.userID)
			req := httptest.NewRequest("DELETE", url, nil)

			// Executar requisição
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verificar status
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedError {
				// Verificar resposta de erro
				var errorResponse map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
				require.NoError(t, err)

				assert.Contains(t, errorResponse, "error")
			}
		})
	}
}
