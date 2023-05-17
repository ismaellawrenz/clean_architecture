package router

import (
	"crud-golang/internal/app/handler"
	"crud-golang/internal/app/repository"
	"crud-golang/internal/app/service"
	"database/sql"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func New() *echo.Echo {

	// Configuração do banco de dados
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	// Inicialização dos pacotes
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()

	e.POST("/users", userHandler.CreateUserHandler)
	e.GET("/users/:id", userHandler.GetUserByIDHandler)
	e.PUT("/users", userHandler.UpdateUserHandler)
	e.DELETE("users/:id", userHandler.DeleteUserByIDHandler)

	err = e.Start(":8080")

	fmt.Println("Server listening on port 8080")

	if err != nil {
		log.Fatal("failed to start server:", err)
	}

	return e
}
