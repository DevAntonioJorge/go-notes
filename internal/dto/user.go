package dto

type CreateUserRequest struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct{
	Identifier string `json:"identifier"`
	Password string `json:"password"`
}

