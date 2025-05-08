package interfaces

import (
	"github.com/DevAntonioJorge/go-blog/internal/dto"
	"github.com/DevAntonioJorge/go-blog/internal/models"
)

type IUserService interface{
	SaveUser(input dto.CreateUserRequest) error
	Login(input dto.LoginRequest, valueType string) (*models.User, error)
	UpdatePassword(password string) error
}