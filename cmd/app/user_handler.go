package main

import "github.com/DevAntonioJorge/go-blog/internal/interfaces"

type UserHandler struct{
	service interfaces.IUserService
}

func NewUserHandler(service interfaces.IUserService) *UserHandler{
	return &UserHandler{service}
}