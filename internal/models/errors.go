package models

import "errors"

var (

	ErrInvalidName = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)