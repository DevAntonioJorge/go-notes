package main

import (
	"log"
	"net/http"

	"github.com/DevAntonioJorge/go-notes/internal/dto"
	"github.com/DevAntonioJorge/go-notes/internal/interfaces"
	"github.com/DevAntonioJorge/go-notes/internal/utils/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service interfaces.IUserService
}

func NewUserHandler(service interfaces.IUserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) RegisterHandler(c echo.Context) error {
	var u dto.CreateUserRequest
	
	if err := c.Bind(&u); err != nil {
		log.Printf("Error binding json: %v", err)
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	validate := validator.New()
	if err:= validate.Struct(u); err != nil{
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}
	
	if err := h.service.SaveUser(u); err != nil {
		log.Printf("Error saving user: %v", err)
		return c.String(http.StatusBadRequest, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User successfully registered"})
}

func (h *UserHandler) LoginHandler(c echo.Context) error {
	var lr dto.LoginRequest
	if err := c.Bind(&lr); err != nil {
		log.Printf("Error binding json: %v", err)
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	user, err := h.service.Login(lr)
	if err != nil {
		log.Printf("Error processing login: %v", err)
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}
	token, err := token.GenerateToken(user.ID, GetConfig().JWTSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating token")
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "User login successful",
		"token":   token,
		"user":    user,
	})
}

func (h *UserHandler) UpdatePasswordHandler(c echo.Context) error {
	var req dto.UpdatePasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadGateway, "Invalid request body")
	}

	user := c.Get("user_id").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	if err := h.service.UpdatePassword(userId, req.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "Password successfully updated"})
}
