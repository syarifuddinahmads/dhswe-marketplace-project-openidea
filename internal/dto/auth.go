package dto

import "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"

type AuthLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	model.User
}

// Register
type AuthRegisterRequest struct {
	model.User
}
type AuthRegisterResponse struct {
	model.User
}
