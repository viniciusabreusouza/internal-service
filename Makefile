.PHONY: test test-unit test-integration test-all build clean deps lint coverage dev-setup check help

# Comandos de teste
test: test-unit test-integration

test-unit:
	@echo "Executando testes unitários..."
	go clean -testcache
	go test -v ./... -tags=unit

test-integration:
	@echo "Executando testes de integração..."
	go clean -testcache
	go test -v ./... -tags=integration

test-all:
	@echo "Executando todos os testes..."
	go clean -testcache
	go test -v ./... -tags="unit integration"

# Comandos de build
build:
	@echo "Compilando o projeto..."
	go build -o server ./cmd/server

# Comandos de limpeza
clean:
	@echo "Limpando arquivos de build..."
	rm -f server
	go clean -cache

clean-test-cache:
	@echo "Limpando cache de testes..."
	go clean -testcache

clean-all: clean clean-test-cache
	@echo "Limpeza completa concluída!"

# Comandos de dependências
deps:
	@echo "Baixando dependências..."
	go mod download
	go mod tidy

# Comandos de linting
lint:
	@echo "Executando linter..."
	golangci-lint run

# Comandos de cobertura
coverage:
	@echo "Gerando relatório de cobertura..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Relatório de cobertura gerado em coverage.html"

# Comandos de desenvolvimento
dev-setup: deps
	@echo "Configurando ambiente de desenvolvimento..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Comandos de verificação
check: lint test
	@echo "Verificações concluídas com sucesso!"

# Ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  test              - Executa testes unitários e de integração"
	@echo "  test-unit         - Executa apenas testes unitários"
	@echo "  test-integration  - Executa apenas testes de integração"
	@echo "  test-all          - Executa todos os testes do projeto"
	@echo "  build             - Compila o projeto"
	@echo "  clean             - Remove arquivos de build"
	@echo "  clean-test-cache  - Remove cache de testes"
	@echo "  clean-all         - Remove todos os caches"
	@echo "  deps              - Baixa e organiza dependências"
	@echo "  lint              - Executa linter"
	@echo "  coverage          - Gera relatório de cobertura"
	@echo "  dev-setup         - Configura ambiente de desenvolvimento"
	@echo "  check             - Executa lint e testes"
	@echo "  help              - Mostra esta ajuda" 