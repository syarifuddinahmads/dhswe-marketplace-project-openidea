package auth

import (
	"context"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/factory"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	res "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/util/response"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

type service struct {
	Repository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository}
}

func (s *service) Login(ctx context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Token: token,
		User:  *data,
	}

	return result, nil
}

func (s *service) Register(ctx context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse

	// Create a model.User object with information received from the payload
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	user := model.User{
		Name:     payload.Name,
		Username: payload.Username,
		Password: payload.Password,
	}

	// Call the Repository method to register the user
	err := s.Repository.Register(ctx, user)
	if err != nil {
		// Handle error and return the appropriate response
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Generate a token for the registered user
	token, err := user.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Create the AuthRegisterResponse using the user and token
	result = &dto.AuthRegisterResponse{
		User:  user,
		Token: token,
	}

	// Return the successful result
	return result, nil

}
