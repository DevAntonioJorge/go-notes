package interfaces

import "github.com/DevAntonioJorge/go-blog/internal/models"

type IUserRepository interface {
	SaveUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
	UpdatePassword(user *models.User, password string) error
}