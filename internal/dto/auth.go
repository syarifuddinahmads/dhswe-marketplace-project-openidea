package dto

import (
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

// Login
type AuthLoginRequest struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	model.User
}

// Register
type AuthRegisterRequest struct {
	model.User
}
type AuthRegisterResponse struct {
	model.User
}
