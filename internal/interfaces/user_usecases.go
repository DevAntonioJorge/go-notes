package interfaces

import (
	"github.com/DevAntonioJorge/go-blog/internal/models"
)

type IUserService interface{
	SaveUser(user *models.User) error
	Login(user *models.User, identifier string, idType string) error
	ValidatePassword(user *models.User, password string) error
}