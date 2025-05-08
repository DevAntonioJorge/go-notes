package interfaces

import (
	"github.com/DevAntonioJorge/go-blog/internal/models"
)

type IUserService interface{
	SaveUser(user *models.User) error
}