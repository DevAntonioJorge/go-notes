package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password []byte `json:"-"`
	CreatedAt string `json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	validate := validator.New()
	if name == ""{
		return nil, ErrInvalidName
	}
	if err := validate.Var(email, "required,email"); err != nil{
		return nil, ErrInvalidEmail
	}
	if password == "" || len(password) < 8 || len(password) > 72{
		return nil, ErrInvalidPassword
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	return &User{
		ID: uuid.NewString(),
		Name: name,
		Email: email,
		Password: hashed,
		CreatedAt: time.Now().Format(time.DateTime),
	},nil
}

func (u *User) CheckPassword(password string) error{
	if err:= bcrypt.CompareHashAndPassword(u.Password, []byte(password)); err != nil{
		return err
	}
	return nil
}