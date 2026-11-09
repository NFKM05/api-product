# api-product

API REST em Go para cadastro de produtos vinculados a usuários, com autenticação via JWT e controle de posse.

## Stack

- **Go** (net/http, sem framework)
- **PostgreSQL**
- **sqlc** — geração de código a partir de SQL puro
- **golang-migrate** — versionamento e aplicação automática de migrations
- **JWT** (`golang-jwt/jwt/v5`) — autenticação
- **bcrypt** — hash de senha
- **Swagger / OpenAPI** (`swaggo`) — documentação interativa da API

## Arquitetura

O projeto segue uma arquitetura em camadas:

```
cmd/api            -> ponto de entrada (main.go)
internal/
  domain/          -> entidades e regras de negocio puras (validacao, normalizacao)
  dto/             -> estruturas de request/response da API
  repository/      -> acesso ao banco (traduz domain <-> sqlc)
  service/         -> regras de negocio, orquestracao, checagem de dono
  handler/         -> camada HTTP (decodifica request, chama service, escreve response)
  middleware/       -> autenticacao JWT e logging de requisicoes
  routes/          -> registro das rotas no http.ServeMux
  bootstrap/       -> injecao de dependencia, monta a aplicacao
  db/              -> codigo gerado pelo sqlc (nao editar manualmente)
  config/          -> leitura de variaveis de ambiente
  connection/      -> conexao com o banco e execucao das migrations
  utils/           -> funcoes auxiliares (resposta de erro, hash, etc.)
migrations/        -> arquivos .up.sql / .down.sql do golang-migrate
sql/
  schema.sql       -> referencia do schema (usado pelo sqlc, nao aplicado ao banco)
  queries/         -> queries anotadas, usadas pelo sqlc para gerar o codigo em internal/db
docs/              -> especificacao Swagger gerada automaticamente
```

Fluxo de uma requisição: `routes` → `middleware` (autenticação) → `handler` (HTTP) → `service` (regra de negócio) → `repository` (banco) → `domain` (entidade).

## Pré-requisitos

- Go 1.26+
- PostgreSQL rodando localmente

## Configuração

1. Clone o repositório e entre na pasta do projeto.
2. Copie o arquivo de exemplo de variáveis de ambiente:
   ```
   cp .env.example .env
   ```
3. Ajuste os valores conforme seu ambiente:
   ```
   DATABASE_URL=postgresql://postgres:postgres@localhost:5432/api_product?sslmode=disable
   MIGRATION_PATH=./migrations
   SERVICE_PORT=8080
   ```
4. Crie o banco de dados vazio (as tabelas são criadas automaticamente na próxima etapa):
   ```sql
   CREATE DATABASE api_product;
   ```

## Rodando o projeto

Na raiz do projeto:
```
go run ./cmd/api
```
Ao subir, a aplicação aplica as migrations automaticamente, cria `users` e `products`, e sobe o servidor na porta definida em `SERVICE_PORT` (padrão `8080`).

## Documentação interativa (Swagger)

Com a aplicação rodando, acesse:
```
http://localhost:8080/swagger/index.html
```
Lá é possível ver todos os endpoints, seus parâmetros, respostas esperadas, e testar diretamente pelo navegador (inclusive os protegidos, usando o botão **Authorize** para informar o token).

Sempre que uma rota ou DTO for alterado, a documentação precisa ser regenerada:
```
swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
```

## Autenticação

1. Cria uma conta: `POST /users`
2. Faz login: `POST /login` — devolve um token JWT (`access_token`)
3. Usa o token nas rotas protegidas, no header:
   ```
   Authorization: Bearer <token>
   ```

Todas as rotas exigem autenticação, exceto `POST /users` (cadastro) e `POST /login`.

Operações de escrita sobre produto (`PUT`, `DELETE`) só são permitidas para o usuário que criou o produto. Operações sobre a própria conta de usuário (`UpdateName`, `UpdateEmail`, `UpdatePassword`, `DesactivateUser`) só são permitidas para o próprio usuário autenticado — tentativas sobre outro usuário/produto retornam `403 Forbidden`.

## Endpoints

### Usuários

| Método | Rota | Autenticação | Descrição |
|---|---|---|---|
| POST | `/users` | não | Cadastra um novo usuário |
| POST | `/login` | não | Autentica e devolve um token JWT |
| GET | `/users` | sim | Lista todos os usuários |
| GET | `/users/{id}` | sim | Busca um usuário por id |
| PATCH | `/users/{id}/change-name` | sim (dono) | Altera o nome |
| PATCH | `/users/{id}/change-email` | sim (dono) | Altera o email |
| PATCH | `/users/{id}/change-password` | sim (dono) | Altera a senha (exige a senha atual) |
| PATCH | `/users/{id}/desactivate` | sim (dono) | Desativa a conta |

### Produtos

| Método | Rota | Autenticação | Descrição |
|---|---|---|---|
| POST | `/products` | sim | Cria um produto (vinculado ao usuário autenticado) |
| GET | `/products` | sim | Lista todos os produtos |
| GET | `/products/{id}` | sim | Busca um produto por id |
| PUT | `/products/{id}` | sim (dono) | Atualiza um produto |
| DELETE | `/products/{id}` | sim (dono) | Exclui um produto |

## Migrations

Gerenciadas pelo `golang-migrate`, aplicadas automaticamente ao subir a aplicação. Os arquivos ficam em `migrations/`, numerados em ordem de aplicação (`0001_...`, `0002_...`). Alterações de schema devem ser feitas criando uma nova migration, não editando uma já aplicada.
