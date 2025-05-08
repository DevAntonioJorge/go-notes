package interfaces

import "github.com/DevAntonioJorge/go-blog/internal/dto"

type IUserService interface{
	SaveUser(inpur dto.CreateUserRequest) error
}