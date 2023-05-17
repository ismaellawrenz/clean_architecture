# CRUD com Clean Architecture em Golang e PostgreSQL

Este é um projeto de exemplo que implementa um CRUD (Create, Read, Update, Delete) utilizando a Clean Architecture em Golang com PostgreSQL como banco de dados.

## Pré-requisitos

Certifique-se de ter os seguintes requisitos instalados em sua máquina:

- Golang (versão X.X.X)
- PostgreSQL (versão X.X.X)

## Configuração do banco de dados

1. Crie um banco de dados PostgreSQL.
2. Configure as credenciais do banco de dados no arquivo `config.toml`.

## Instalação

1. Clone este repositório:
```bash
git clone https://github.com/seu-usuario/nome-do-repositorio.git
```

2. Instale as dependências do projeto:
```bash 
go mod tidy
```

3. Execute o aplicativo:
```bash 
go run main.go
```

## API Endpoints

A seguir estão os endpoints da API disponíveis:

- `GET /users/{id}`: Retorna um usuário pelo ID.
- `POST /users`: Cria um novo usuário.
- `PUT /users`: Atualiza um usuário existente.
- `DELETE /users/{id}`: Exclui um usuário pelo ID.
