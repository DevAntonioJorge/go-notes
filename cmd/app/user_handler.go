package main

import (
	"net/http"

	"github.com/DevAntonioJorge/go-blog/internal/dto"
	"github.com/DevAntonioJorge/go-blog/internal/interfaces"
	"github.com/DevAntonioJorge/go-blog/internal/models"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	service interfaces.IUserService
}

func NewUserHandler(service interfaces.IUserService) *UserHandler{
	return &UserHandler{service}
}

func (h *UserHandler) RegisterHandler(c echo.Context) error{
	u := new(dto.CreateUserRequest)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	user, err := models.NewUser(u.Name, u.Email, u.Password)
	if err != nil{
		return c.String(http.StatusBadRequest, "Invalid name, email or password")
	}
	if err := h.service.SaveUser(user); err != nil{
		return c.String(http.StatusBadRequest, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User successfully registered", "user": user})
}

func (h *UserHandler) LoginHandler(c echo.Context) error{
	
	return c.JSON(http.StatusAccepted, "User login successful")
}