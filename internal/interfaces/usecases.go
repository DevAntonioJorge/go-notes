package interfaces

import (
	"github.com/DevAntonioJorge/go-notes/internal/dto"
	"github.com/DevAntonioJorge/go-notes/internal/models"
)

type IUserService interface{
	SaveUser(input dto.CreateUserRequest) error
	Login(input dto.LoginRequest) (*models.User, error)
	UpdatePassword(id , password string) error
}