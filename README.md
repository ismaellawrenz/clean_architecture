# CRUD com Clean Architecture em Golang e PostgreSQL

Este é um projeto de exemplo que implementa um CRUD (Create, Read, Update, Delete) utilizando a Clean Architecture em Golang com PostgreSQL como banco de dados.

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
