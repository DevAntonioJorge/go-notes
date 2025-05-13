package models

import "errors"

var (

	ErrInvalidName = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidFields = errors.New("invalid fields")
	ErrUserNotFound = errors.New("user not found")
	ErrUpdatePassword = errors.New("failed to update password")
)