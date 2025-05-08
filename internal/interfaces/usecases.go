package interfaces

import (
	"github.com/DevAntonioJorge/go-blog/internal/models"
)

type IUserService interface{
	SaveUser(user *models.User) error
	Login(identifier string, idType, password string) (*models.User, error)
	UpdatePassword(password string) error
}