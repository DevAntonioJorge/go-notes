package api

import (
	"log"
	"net/http"

	"github.com/DevAntonioJorge/go-notes/internal/domain/utils/token"
	"github.com/DevAntonioJorge/go-notes/internal/infra/config"
	"github.com/DevAntonioJorge/go-notes/internal/infra/dto"
	"github.com/DevAntonioJorge/go-notes/pkg/schema"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterHandler(c echo.Context) error {
	var u dto.CreateUserRequest

	if err := c.Bind(&u); err != nil {
		log.Printf("Error binding json: %v", err)
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	if err := schema.NewValidator().Validate(&u); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}

	if err := s.service.User.SaveUser(u); err != nil {
		log.Printf("Error saving user: %v", err)
		return c.String(http.StatusBadRequest, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User successfully registered"})
}

func (s *Server) LoginHandler(c echo.Context) error {
	var lr dto.LoginRequest
	if err := c.Bind(&lr); err != nil {
		log.Printf("Error binding json: %v", err)
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	if err := schema.NewValidator().Validate(&lr); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}
	user, err := s.service.User.Login(lr)
	if err != nil {
		log.Printf("Error processing login: %v", err)
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}
	token, err := token.GenerateToken(user.ID, config.GetConfig().JWTSecret)
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

func (s *Server) UpdatePasswordHandler(c echo.Context) error {
	var req dto.UpdatePasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadGateway, "Invalid request body")
	}
	if err := schema.NewValidator().Validate(&req); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}
	user := c.Get("user_id").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	if err := s.service.User.UpdatePassword(userId, req.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "Password successfully updated"})
}
