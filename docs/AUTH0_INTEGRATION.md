# Integração com Auth0 - Proteção de Rotas JWT

Esta documentação descreve a implementação da proteção de rotas usando tokens JWT do Auth0.

## Configuração

### Variáveis de Ambiente

A aplicação suporta as seguintes variáveis de ambiente para configuração do Auth0:

- `AUTH0_DOMAIN`: Domínio do Auth0 (padrão: `dev-minhaempresa.us.auth0.com`)
- `AUTH0_AUDIENCE`: Audience da API (padrão: `https://minha-api.com`)

### Configuração no Auth0

1. **Criar uma API no Auth0:**
   - Acesse o Dashboard do Auth0
   - Vá para "Applications" > "APIs"
   - Clique em "Create API"
   - Configure:
     - Name: `Minha API`
     - Identifier: `https://minha-api.com`
     - Signing Algorithm: `RS256`

2. **Configurar Client Credentials Flow:**
   - Vá para "Applications" > "Applications"
   - Crie uma nova aplicação do tipo "Machine to Machine"
   - Configure as permissões necessárias

## Rotas Protegidas

### GET /dados-protegidos

Esta rota requer autenticação via JWT token.

**Headers necessários:**
```
Authorization: Bearer <seu-jwt-token>
```

**Resposta de sucesso (200 OK):**
```json
{
  "message": "Token válido! Acesso permitido.",
  "token_info": {
    "subject": "client-id@clients",
    "audience": ["https://minha-api.com"],
    "issuer": "https://dev-minhaempresa.us.auth0.com/",
    "issued_at": "2024-01-15T10:30:00Z",
    "expires_at": "2024-01-15T11:30:00Z"
  },
  "user_id": "client-id@clients"
}
```

**Resposta de erro (401 Unauthorized):**
```json
{
  "error": "Authorization header is required"
}
```

ou

```json
{
  "error": "Invalid token: failed to parse/validate token: token is expired"
}
```

## Validação do Token

O middleware de autenticação valida os seguintes aspectos do token JWT:

1. **Assinatura**: Verifica a assinatura usando as chaves públicas do Auth0 (JWKS)
2. **Audience**: Confirma que o token foi emitido para `https://minha-api.com`
3. **Issuer**: Verifica que o token foi emitido por `https://dev-minhaempresa.us.auth0.com/`
4. **Expiração**: Verifica se o token não expirou
5. **Formato**: Valida o formato do header Authorization

## Exemplo de Uso

### Testando com curl

```bash
# Substitua <SEU_TOKEN_JWT> pelo token real obtido do Auth0
curl -X GET http://localhost:8080/dados-protegidos \
  -H "Authorization: Bearer <SEU_TOKEN_JWT>"
```

### Obtendo um token via Client Credentials Flow

```bash
# Substitua os valores pelos seus dados reais
curl -X POST https://dev-minhaempresa.us.auth0.com/oauth/token \
  -H "Content-Type: application/json" \
  -d '{
    "client_id": "SEU_CLIENT_ID",
    "client_secret": "SEU_CLIENT_SECRET",
    "audience": "https://minha-api.com",
    "grant_type": "client_credentials"
  }'
```

## Implementação Técnica

### Estrutura de Arquivos

```
internal/
├── config/
│   └── auth.go          # Configurações do Auth0
├── middleware/
│   └── auth.go          # Middleware de autenticação JWT
└── handler/
    └── protected_handler.go  # Handler da rota protegida
```

### Dependências

- `github.com/lestrrat-go/jwx/v2/jwk`: Para carregar chaves públicas do JWKS
- `github.com/lestrrat-go/jwx/v2/jwt`: Para validação de tokens JWT
- `github.com/gin-gonic/gin`: Framework web

### Funcionalidades

- ✅ Validação de assinatura RS256
- ✅ Verificação de audience
- ✅ Verificação de issuer
- ✅ Verificação de expiração
- ✅ Configuração via variáveis de ambiente
- ✅ Logs detalhados
- ✅ Respostas de erro claras
- ✅ Tolerância de 30 segundos para diferenças de relógio

## Troubleshooting

### Erro: "failed to fetch JWKS"
- Verifique se o domínio do Auth0 está correto
- Confirme se o endpoint JWKS está acessível

### Erro: "token is expired"
- O token expirou, obtenha um novo token

### Erro: "invalid audience"
- Verifique se o audience no token corresponde ao configurado
- Confirme se o token foi emitido para a API correta

### Erro: "invalid issuer"
- Verifique se o issuer no token corresponde ao domínio do Auth0 configurado 