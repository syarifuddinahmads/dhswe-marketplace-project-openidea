package dto

import (
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

// Login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	model.User
}

// Register
type AuthRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthRegisterResponse struct {
	model.User
	Token string `json:"token"`
}
