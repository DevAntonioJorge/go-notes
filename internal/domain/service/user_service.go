package service

import (
	"strings"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"github.com/DevAntonioJorge/go-notes/internal/domain/repository"
	"github.com/DevAntonioJorge/go-notes/internal/infra/dto"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) SaveUser(u dto.CreateUserRequest) error {
	_, err := s.repo.User.GetUserByEmail(u.Email)
	if err == nil {
		return models.ErrEmailExists
	}
	user := models.NewUser(u.Name, u.Email, u.Password)
	if err = s.repo.User.SaveUser(user); err != nil {
		return err
	}
	return nil
}
func (s *UserService) Login(input dto.LoginRequest) (*models.User, error) {
	var user *models.User
	var err error

	if strings.Contains(input.Identifier, "@") {
		user, err = s.repo.User.GetUserByEmail(input.Identifier)
	} else {
		user, err = s.repo.User.GetUserByName(input.Identifier)
	}

	if user == nil || err != nil {
		return nil, models.ErrUserNotFound
	}

	if err = user.CheckPassword(input.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdatePassword(id, password string) error {

	user, err := s.repo.User.GetUserByID(id)
	if err != nil {
		return models.ErrUserNotFound
	}
	if err := s.repo.User.UpdatePassword(user, password); err != nil {
		return models.ErrUpdatePassword
	}

	return nil
}
