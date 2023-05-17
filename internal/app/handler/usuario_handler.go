package handler

import (
	"crud-golang/internal/app/model"
	"crud-golang/internal/app/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUserHandler(c echo.Context) error {

	var user model.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		http.Error(c.Response(), "invalid request payload", http.StatusBadRequest)
		return nil
	}

	userNew, err := h.userService.CreateUser(&user)
	userNew.Password = ""
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		return nil
	}

	return c.JSON(http.StatusOK, userNew)
}

func (h *UserHandler) GetUserByIDHandler(c echo.Context) error {

	idParam := c.Param("id")

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "ID invalido")
	}

	user, err := h.userService.GetUserByID(int64(userID))
	if err != nil {

		return c.JSON(http.StatusInternalServerError, "Usuário não encontrado")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUserHandler(c echo.Context) error {

	var user model.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		http.Error(c.Response(), "invalid request payload", http.StatusBadRequest)
		return nil
	}

	userNew, err := h.userService.Update(&user)

	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		return nil
	}
	userNew.Password = ""
	return c.JSON(http.StatusOK, userNew)
}

func (h *UserHandler) DeleteUserByIDHandler(c echo.Context) error {

	idParam := c.Param("id")

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "ID invalido")
	}

	err = h.userService.Delete(int64(userID))
	if err != nil {

		return c.JSON(http.StatusInternalServerError, "Falha ao deletar usuário"+err.Error())
	}

	return c.NoContent(http.StatusOK)
}
