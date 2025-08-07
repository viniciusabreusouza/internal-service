package integration

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	sharedContainer testcontainers.Container
	sharedClient    *mongo.Client
	containerOnce   sync.Once
	containerMutex  sync.Mutex
)

// StartSharedMongoContainer inicia um container MongoDB compartilhado
func StartSharedMongoContainer(t *testing.T) (*mongo.Client, func()) {
	containerMutex.Lock()
	defer containerMutex.Unlock()

	// Se já existe um container, retorna o cliente existente
	if sharedContainer != nil && sharedClient != nil {
		return sharedClient, func() {
			// Cleanup será feito no final de todos os testes
		}
	}

	ctx := context.Background()
	mongoContainer, err := mongodb.RunContainer(ctx,
		testcontainers.WithImage("mongo:7.0"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("Waiting for connections"),
		),
	)
	if err != nil {
		t.Fatalf("failed to start shared container: %v", err)
	}

	mongoURI, err := mongoContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("failed to ping MongoDB: %v", err)
	}

	// Armazena referências globais
	sharedContainer = mongoContainer
	sharedClient = client

	// Retorna função de cleanup que será chamada no final
	cleanup := func() {
		// Cleanup será feito por CleanupSharedMongoContainer
	}

	return client, cleanup
}

// CleanupSharedMongoContainer limpa o container MongoDB compartilhado
func CleanupSharedMongoContainer(t *testing.T) {
	containerMutex.Lock()
	defer containerMutex.Unlock()

	if sharedClient != nil {
		ctx := context.Background()
		if err := sharedClient.Disconnect(ctx); err != nil {
			t.Logf("failed to disconnect from MongoDB: %v", err)
		}
		sharedClient = nil
	}

	if sharedContainer != nil {
		ctx := context.Background()
		if err := sharedContainer.Terminate(ctx); err != nil {
			t.Logf("failed to terminate shared container: %v", err)
		}
		sharedContainer = nil
	}
}

// StartMongoContainer inicia um container MongoDB para testes
func StartMongoContainer(t *testing.T) (testcontainers.Container, *mongo.Client) {
	ctx := context.Background()
	mongoContainer, err := mongodb.RunContainer(ctx,
		testcontainers.WithImage("mongo:7.0"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("Waiting for connections"),
		),
	)
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}
	mongoURI, err := mongoContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("failed to connect to MongoDB: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("failed to ping MongoDB: %v", err)
	}
	return mongoContainer, client
}

// CleanupMongoContainer limpa o container MongoDB
func CleanupMongoContainer(t *testing.T, container testcontainers.Container, client *mongo.Client) {
	ctx := context.Background()
	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			t.Logf("failed to disconnect from MongoDB: %v", err)
		}
	}
	if container != nil {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %v", err)
		}
	}
}

// GetTestDatabase retorna uma instância do banco de dados para testes
func GetTestDatabase(client *mongo.Client) *mongo.Database {
	return client.Database("test_db")
}

// CleanupTestDatabase limpa todas as collections do banco de teste
func CleanupTestDatabase(client *mongo.Client) error {
	ctx := context.Background()
	db := client.Database("test_db")

	// Lista todas as collections
	collections, err := db.ListCollectionNames(ctx, map[string]interface{}{})
	if err != nil {
		return err
	}

	// Remove todas as collections
	for _, collectionName := range collections {
		if err := db.Collection(collectionName).Drop(ctx); err != nil {
			return err
		}
	}

	return nil
}
