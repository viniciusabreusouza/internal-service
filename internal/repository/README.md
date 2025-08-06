# Repositório MongoDB

Este diretório contém a implementação do repositório MongoDB para a entidade User, incluindo testes unitários e de integração.

## Estrutura de Arquivos

- `user_repository.go` - Interface UserRepository
- `user_mongo_repository.go` - Implementação MongoDB do UserRepository
- `user_mongo_repository_test.go` - Testes unitários
- `user_mongo_repository_integration_test.go` - Testes de integração
- `user_memory_repository.go` - Implementação em memória (existente)

## Implementação MongoDB

### Características

- **Interface CollectionInterface**: Permite mock para testes unitários
- **Conversão de tipos**: Funções `toDocument()` e `toUser()` para conversão entre domínio e documento MongoDB
- **Paginação**: Suporte completo para paginação no método `GetAll()`
- **Tratamento de erros**: Tratamento adequado de erros do MongoDB
- **Validação**: Verificação de existência de registros antes de operações

### Métodos Implementados

1. **Create**: Cria um novo usuário no MongoDB
2. **GetByID**: Busca um usuário pelo ID
3. **GetAll**: Lista usuários com paginação
4. **Update**: Atualiza um usuário existente
5. **Delete**: Remove um usuário pelo ID

## Testes

### Testes Unitários

- Usam mocks para simular o comportamento do MongoDB
- Testam cenários de sucesso e falha
- Cobertura dos métodos principais

### Testes de Integração

- Usam Testcontainers para criar containers MongoDB reais
- Testam a integração completa com MongoDB
- Cobertura de cenários CRUD completos

## Execução dos Testes

### Usando Tags (Recomendado)

```bash
# Testes unitários
make test-unit
make test-unit-fresh        # Sem cache

# Testes de integração
make test-integration
make test-integration-fresh # Sem cache

# Todos os testes
make test

# Comandos específicos do MongoDB
make test-mongo-unit
make test-mongo-unit-fresh
make test-mongo-integration
make test-mongo-integration-fresh

# Limpeza de cache
make clean-test-cache
make clean-all
```

### Execução Direta com Tags

```bash
# Testes unitários
go test -v ./internal/repository -tags=unit

# Testes de integração
go test -v ./internal/repository -tags=integration

# Todos os testes (sem tags)
go test -v ./internal/repository
```

## Dependências

- `go.mongodb.org/mongo-driver` - Driver oficial do MongoDB
- `github.com/testcontainers/testcontainers-go` - Para testes de integração
- `github.com/stretchr/testify` - Para assertions e mocks

## Configuração

O repositório MongoDB requer uma conexão com MongoDB. Para testes, os containers são criados automaticamente usando Testcontainers.

### Exemplo de Uso

```go
// Conectar ao MongoDB
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
    log.Fatal(err)
}

// Criar repositório
db := client.Database("myapp")
repo := NewUserMongoRepository(db)

// Usar o repositório
user, err := user.NewUser("John Doe", "john@example.com")
if err != nil {
    log.Fatal(err)
}

err = repo.Create(ctx, user)
if err != nil {
    log.Fatal(err)
}
```

## Organização dos Testes

### Build Tags

Os testes usam build tags para separação:

- **`//go:build unit`**: Testes unitários
- **`//go:build integration`**: Testes de integração

### Testes Unitários (`user_mongo_repository_test.go`)

- Testam a lógica do repositório isoladamente
- Usam mocks para simular o MongoDB
- Focam em cenários específicos de cada método
- Tag: `unit`

### Testes de Integração (`user_mongo_repository_integration_test.go`)

- Testam a integração real com MongoDB
- Usam containers Docker via Testcontainers
- Cobertura de cenários end-to-end
- Tag: `integration`

### Código Comum (`tests/integration/common.go`)

- Funções reutilizáveis para inicialização de containers
- Configuração padrão para testes de integração
- Limpeza automática de recursos

## Padrões Utilizados

1. **Table Driven Tests**: Para cenários múltiplos
2. **Interface Segregation**: CollectionInterface para testabilidade
3. **Dependency Injection**: Injeção de dependências via construtor
4. **Error Handling**: Tratamento adequado de erros
5. **Clean Architecture**: Separação clara entre domínio e infraestrutura 