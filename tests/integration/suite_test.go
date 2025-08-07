//go:build integration
// +build integration

package integration

import (
	"os"
	"testing"
)

// TestMain é executado uma vez antes de todos os testes de integração
func TestMain(m *testing.M) {
	// Executa todos os testes
	code := m.Run()

	// Cleanup do container compartilhado após todos os testes
	CleanupSharedMongoContainer(&testing.T{})

	// Retorna o código de saída
	os.Exit(code)
} 