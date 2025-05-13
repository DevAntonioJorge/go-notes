package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID string `json:"id"`
	Name string `json:"name" validate:"required,min=5,max=20"`
	Email string `json:"email" validate:"required,unique,email"`
	Password []byte `json:"-" validate:"required,min=8,max=72"`
	CreatedAt string `json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	var user *User
	validate := validator.New()
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil, ErrInvalidPassword
	}
	user = &User{
		ID: uuid.NewString(),
		Name: name,
		Email: email,
		Password: hashed,
		CreatedAt: time.Now().Format(time.DateTime),
	}
	if err = validate.Struct(user); err != nil{
		return nil, ErrInvalidFields
	}
	return user, nil
}

func (u *User) CheckPassword(password string) error{
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(password)); err != nil{
		return err
	}
	return nil
}