package main

import (
	"net/http"

	"github.com/DevAntonioJorge/go-blog/internal/dto"
	"github.com/DevAntonioJorge/go-blog/internal/interfaces"
	"github.com/DevAntonioJorge/go-blog/internal/utils/token"
	"github.com/golang-jwt/jwt/v5"
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
	var u dto.CreateUserRequest
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	if err := h.service.SaveUser(u); err != nil{
		return c.String(http.StatusBadRequest, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User successfully registered"})
}

func (h *UserHandler) LoginHandler(c echo.Context) error{
	var lr dto.LoginRequest
	if err := c.Bind(lr); err != nil{
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	user, err := h.service.Login(lr)
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
		"user": user,
	})
}

func (h *UserHandler) UpdatePasswordHandler(c echo.Context) error{
	var req struct {
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil{
		return c.String(http.StatusBadGateway, "Invalid request body")
	}

	user := c.Get("user_id").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	if err := h.service.UpdatePassword(userId, req.Password); err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message":"Password successfully updated"})
}
