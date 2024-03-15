package auth

import (
	"context"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/response"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s *Service) Login(ctx context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.repo.FindByUsername(ctx, payload.Username)

	if data == nil {
		return result, response.ErrorBuilder(&response.ErrorConstant.UsernameOrPasswordIncorrect, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		return result, response.ErrorBuilder(&response.ErrorConstant.UsernameOrPasswordIncorrect, err)
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		AccessToken: token,
		Username:    data.Username,
		Name:        data.Name,
	}

	return result, nil
}

func (s *Service) Register(ctx context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	// Replace the plain text password with the hashed password
	payload.Password = string(hashedPassword)

	// Call the repository Register function
	data, err := s.repo.Register(ctx, payload)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	token, err := data.GenerateToken()
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result := &dto.AuthRegisterResponse{
		AccessToken: token,
		Username:    payload.Username,
		Name:        payload.Name,
	}

	return result, nil
}
