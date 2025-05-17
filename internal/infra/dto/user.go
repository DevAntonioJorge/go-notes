package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=5,max=20"`
	Email    string `json:"email" validate:"required,unique,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required,min=5,max=20"`
	Password   string `json:"password" validate:"required,min=8,max=72"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,min=8,max=72"`
}
