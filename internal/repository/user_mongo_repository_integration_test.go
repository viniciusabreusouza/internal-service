//go:build integration
// +build integration

package repository

import (
	"context"
	"testing"

	"example.com/internal-service/internal/domain/user"
	"example.com/internal-service/tests/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserMongoRepository_Integration_Create(t *testing.T) {
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Limpar banco antes do teste
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	repo := NewUserMongoRepository(db)

	ctx := context.Background()

	// Teste de criação bem-sucedida
	testUser, err := user.NewUser("Test User", "test@example.com")
	require.NoError(t, err)

	err = repo.Create(ctx, testUser)
	assert.NoError(t, err)

	// Verifica se o usuário foi criado
	createdUser, err := repo.GetByID(ctx, testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, testUser.ID, createdUser.ID)
	assert.Equal(t, testUser.Name, createdUser.Name)
	assert.Equal(t, testUser.Email, createdUser.Email)
}

func TestUserMongoRepository_Integration_GetByID(t *testing.T) {
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Limpar banco antes do teste
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	repo := NewUserMongoRepository(db)

	ctx := context.Background()

	// Cria um usuário para teste
	testUser, err := user.NewUser("Test User", "test@example.com")
	require.NoError(t, err)

	err = repo.Create(ctx, testUser)
	require.NoError(t, err)

	// Teste de busca bem-sucedida
	foundUser, err := repo.GetByID(ctx, testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, testUser.ID, foundUser.ID)
	assert.Equal(t, testUser.Name, foundUser.Name)
	assert.Equal(t, testUser.Email, foundUser.Email)

	// Teste de usuário não encontrado
	_, err = repo.GetByID(ctx, "non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}

func TestUserMongoRepository_Integration_Update(t *testing.T) {
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Limpar banco antes do teste
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	repo := NewUserMongoRepository(db)

	ctx := context.Background()

	// Cria um usuário para teste
	testUser, err := user.NewUser("Original User", "original@example.com")
	require.NoError(t, err)

	err = repo.Create(ctx, testUser)
	require.NoError(t, err)

	// Teste de atualização bem-sucedida
	err = testUser.Update("Updated User", "updated@example.com")
	require.NoError(t, err)

	err = repo.Update(ctx, testUser)
	assert.NoError(t, err)

	// Verifica se o usuário foi atualizado
	updatedUser, err := repo.GetByID(ctx, testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)

	// Teste de atualização de usuário inexistente
	nonExistentUser := &user.User{
		ID:    "non-existent-id",
		Name:  "Test",
		Email: "test@example.com",
	}
	err = repo.Update(ctx, nonExistentUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}

func TestUserMongoRepository_Integration_Delete(t *testing.T) {
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Limpar banco antes do teste
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	repo := NewUserMongoRepository(db)

	ctx := context.Background()

	// Cria um usuário para teste
	testUser, err := user.NewUser("Test User", "test@example.com")
	require.NoError(t, err)

	err = repo.Create(ctx, testUser)
	require.NoError(t, err)

	// Teste de exclusão bem-sucedida
	err = repo.Delete(ctx, testUser.ID)
	assert.NoError(t, err)

	// Verifica se o usuário foi excluído
	_, err = repo.GetByID(ctx, testUser.ID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	// Teste de exclusão de usuário inexistente
	err = repo.Delete(ctx, "non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}

func TestUserMongoRepository_Integration_GetAll(t *testing.T) {
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Limpar banco antes do teste
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	repo := NewUserMongoRepository(db)

	ctx := context.Background()

	// Cria múltiplos usuários para teste
	users := []*user.User{}
	for i := 1; i <= 5; i++ {
		testUser, err := user.NewUser("User "+string(rune(i+'0')), "user"+string(rune(i+'0'))+"@example.com")
		require.NoError(t, err)
		err = repo.Create(ctx, testUser)
		require.NoError(t, err)
		users = append(users, testUser)
	}

	// Teste de busca com paginação
	foundUsers, total, err := repo.GetAll(ctx, 1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 5, total)
	assert.Len(t, foundUsers, 3)

	// Teste de segunda página
	foundUsers, total, err = repo.GetAll(ctx, 2, 3)
	assert.NoError(t, err)
	assert.Equal(t, 5, total)
	assert.Len(t, foundUsers, 2)

	// Teste com parâmetros inválidos
	foundUsers, total, err = repo.GetAll(ctx, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, 5, total)
	assert.Len(t, foundUsers, 5) // Deve retornar todos os usuários
}

func TestUserMongoRepository_Integration_CRUD_Operations(t *testing.T) {
	client, cleanup := integration.StartSharedMongoContainer(t)
	defer cleanup()

	// Limpar banco antes do teste
	err := integration.CleanupTestDatabase(client)
	require.NoError(t, err)

	db := integration.GetTestDatabase(client)
	repo := NewUserMongoRepository(db)

	ctx := context.Background()

	// Teste completo de CRUD
	// 1. Create
	testUser, err := user.NewUser("CRUD Test User", "crud@example.com")
	require.NoError(t, err)

	err = repo.Create(ctx, testUser)
	assert.NoError(t, err)

	// 2. Read
	foundUser, err := repo.GetByID(ctx, testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, testUser.ID, foundUser.ID)
	assert.Equal(t, testUser.Name, foundUser.Name)
	assert.Equal(t, testUser.Email, foundUser.Email)

	// 3. Update
	err = testUser.Update("Updated CRUD User", "updated-crud@example.com")
	require.NoError(t, err)

	err = repo.Update(ctx, testUser)
	assert.NoError(t, err)

	updatedUser, err := repo.GetByID(ctx, testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated CRUD User", updatedUser.Name)
	assert.Equal(t, "updated-crud@example.com", updatedUser.Email)

	// 4. Delete
	err = repo.Delete(ctx, testUser.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(ctx, testUser.ID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}
