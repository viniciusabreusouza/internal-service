# Dependency Injection com Wire

Este diretório contém a configuração de dependency injection usando o Google Wire.

## Arquivos

- `application.go` - Define a struct `Application` e seus métodos
- `wire_inject.go` - Define os providers e a função de setup do Wire
- `wire_gen.go` - Arquivo gerado automaticamente pelo Wire (não editar)
- `providers.go` - Exemplos de como adicionar novos providers

## Como usar

### 1. Setup básico

```go
import "example.com/internal-service/internal/di"

func main() {
    app, err := di.SetupApplication()
    if err != nil {
        log.Fatal(err)
    }
    
    // Use a aplicação
    logger := app.GetLogger()
    server := app.GetServer()
}
```

### 2. Adicionando novos providers

Para adicionar novos serviços ou repositórios:

1. Crie o provider no arquivo `wire_inject.go`:

```go
func ProvideUserService(userRepo repository.UserRepository) service.UserService {
    return service.NewUserService(userRepo)
}

func ProvideUserRepository(db *sql.DB) repository.UserRepository {
    return repository.NewUserRepository(db)
}
```

2. Adicione o provider ao `wire.Build()`:

```go
func SetupApplication() (Application, error) {
    panic(
        wire.Build(
            ProvideContext,
            ProvideLogger,
            ProvideHTTPServer,
            ProvideUserService,    // Novo provider
            ProvideUserRepository, // Novo provider
            NewApplication,
        ),
    )
}
```

3. Regenerar o código:

```bash
go generate ./internal/di
```

### 3. Regenerando o código

Sempre que você modificar os providers, execute:

```bash
go generate ./internal/di
```

Ou diretamente:

```bash
wire ./internal/di
```

## Estrutura recomendada

```
internal/
├── di/           # Dependency injection
├── domain/       # Entidades e interfaces
├── model/        # Modelos de dados
├── repository/   # Implementações de repositório
└── service/      # Lógica de negócio
```

## Boas práticas

1. **Providers simples**: Mantenha os providers simples e focados
2. **Injeção de dependências**: Use interfaces para desacoplamento
3. **Error handling**: Sempre retorne erros quando apropriado
4. **Testes**: Teste os providers individualmente
5. **Documentação**: Documente providers complexos

## Exemplo completo

Veja `cmd/server/main.go` para um exemplo completo de uso. 