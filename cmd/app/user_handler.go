package main

import (
	"net/http"
	"strings"

	"github.com/DevAntonioJorge/go-blog/internal/dto"
	"github.com/DevAntonioJorge/go-blog/internal/interfaces"
	"github.com/DevAntonioJorge/go-blog/internal/models"
	"github.com/DevAntonioJorge/go-blog/internal/utils/token"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	service interfaces.IUserService
	secret string
}

func NewUserHandler(service interfaces.IUserService, secret string) *UserHandler{
	return &UserHandler{service, secret}
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
	lr := new(dto.LoginRequest)
	var user *models.User
	var err error

	if err := c.Bind(lr); err != nil{
		return c.String(http.StatusBadRequest, "Invalid request body")
	} 
	if strings.Contains(lr.Identifier, "@"){
		user, err = h.service.Login(lr.Identifier, "email", lr.Password)
	} else {
		user, err = h.service.Login(lr.Identifier, "name", lr.Password)
	}
	if err != nil{
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}
	token , err := token.GenerateToken(user.ID, h.secret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating token")
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "User login successful",
		"token": token,
	})
}