package interfaces

import (
	"github.com/DevAntonioJorge/go-blog/internal/dto"
	"github.com/DevAntonioJorge/go-blog/internal/models"
)

type IUserService interface{
	SaveUser(input dto.CreateUserRequest) error
	Login(input dto.LoginRequest) (*models.User, error)
	UpdatePassword(id , password string) error
}