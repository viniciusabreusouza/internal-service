package integration

import (
	"context"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StartMongoContainer inicia um container MongoDB para testes
func StartMongoContainer(t *testing.T) (testcontainers.Container, *mongo.Client) {
	ctx := context.Background()

	// Configura o container MongoDB
	mongoContainer, err := mongodb.RunContainer(ctx,
		testcontainers.WithImage("mongo:7.0"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("Waiting for connections"),
		),
	)
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}

	// Obtém a URI de conexão
	mongoURI, err := mongoContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	// Conecta ao MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Testa a conexão
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
