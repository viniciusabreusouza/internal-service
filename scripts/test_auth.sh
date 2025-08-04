#!/bin/bash

# Script de teste para a autenticação JWT do Auth0
# Este script demonstra como testar a rota protegida

echo "=== Teste de Autenticação JWT Auth0 ==="
echo ""

# URL base da API
API_URL="http://localhost:8080"

echo "1. Testando rota sem token (deve retornar 401):"
echo "curl -X GET $API_URL/dados-protegidos"
echo ""
curl -X GET "$API_URL/dados-protegidos" -w "\nHTTP Status: %{http_code}\n\n"

echo "2. Testando rota com token inválido (deve retornar 401):"
echo "curl -X GET $API_URL/dados-protegidos -H 'Authorization: Bearer invalid-token'"
echo ""
curl -X GET "$API_URL/dados-protegidos" \
  -H "Authorization: Bearer invalid-token" \
  -w "\nHTTP Status: %{http_code}\n\n"

echo "3. Testando rota com formato de header inválido (deve retornar 401):"
echo "curl -X GET $API_URL/dados-protegidos -H 'Authorization: invalid-format'"
echo ""
curl -X GET "$API_URL/dados-protegidos" \
  -H "Authorization: invalid-format" \
  -w "\nHTTP Status: %{http_code}\n\n"

echo "4. Para testar com token válido, use:"
echo "curl -X GET $API_URL/dados-protegidos \\"
echo "  -H 'Authorization: Bearer <SEU_TOKEN_JWT>'"
echo ""
echo "Onde <SEU_TOKEN_JWT> é um token válido obtido do Auth0 via Client Credentials Flow."
echo ""
echo "5. Para obter um token do Auth0, use:"
echo "curl -X POST https://dev-minhaempresa.us.auth0.com/oauth/token \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -d '{"
echo "    \"client_id\": \"SEU_CLIENT_ID\","
echo "    \"client_secret\": \"SEU_CLIENT_SECRET\","
echo "    \"audience\": \"https://minha-api.com\","
echo "    \"grant_type\": \"client_credentials\""
echo "  }'"
echo ""
echo "=== Fim do teste ===" 