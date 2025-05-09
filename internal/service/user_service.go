package service

import (
	"errors"
	"strings"

	"github.com/DevAntonioJorge/go-blog/internal/dto"
	"github.com/DevAntonioJorge/go-blog/internal/interfaces"
	"github.com/DevAntonioJorge/go-blog/internal/models"
)

type UserService struct{
	repo interfaces.IUserRepository
}

func NewUserService(repository interfaces.IUserRepository) *UserService{
	return &UserService{repository}
}
func (s *UserService) SaveUser(u *dto.CreateUserRequest) error {
	_, err := s.repo.GetUserByEmail(u.Email)
	if err == nil{
		return errors.New("user with this email exists")
	}
	_, err = s.repo.GetUserByName(u.Name)
	if err == nil{
		return errors.New("user with this username already exists")
	}
	user, err := models.NewUser(u.Name, u.Email, u.Password)
	if err != nil{
		return errors.New("error saving user")
	}

	if err = s.repo.SaveUser(user); err != nil{
		return err
	}
	return nil
}
func (s *UserService) Login(input dto.LoginRequest) (*models.User, error){
	var user *models.User
	var err error

	if strings.Contains(input.Identifier, "@")  {
		user, err = s.repo.GetUserByEmail(input.Identifier)
	} else {
		user, err = s.repo.GetUserByName(input.Identifier)
	}

	if user == nil || err != nil{
		return nil, errors.New("user not found")
	}
	
	if err = user.CheckPassword(input.Password); err != nil{
		return nil, err
	}

	return user, nil
}
