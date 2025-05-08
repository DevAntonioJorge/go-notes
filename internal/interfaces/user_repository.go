package interfaces

import "github.com/DevAntonioJorge/go-blog/internal/models"

type IUserRepository interface {
	SaveUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdatePassword(email, password string) error
}