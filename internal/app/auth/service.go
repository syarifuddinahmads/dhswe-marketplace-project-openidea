package auth

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
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

func (s *Service) RegisterUser(ctx context.Context, params dto.CreateUserParams) (int, error) {
	// Validate input parameters
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, utils.ErrArgument{Wrapped: err}
	}

	// Check password length
	if len(params.Password) < 5 || len(params.Password) > 15 {
		return 0, errors.New("password length must be between 5 and 15 characters")
	}

	// Start a transaction
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		// Rollback the transaction if there's an error and it hasn't been committed
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Truncate hashed password to 15 characters
	truncatedHash := string(hashedPassword)[:15]

	// Create a User entity with the hashed password
	entity := model.User{
		Name:     params.Name,
		Username: params.Username,
		Password: truncatedHash,
	}

	// Call the repository method to create the user
	err = s.repo.Register(ctx, &entity)
	if err != nil {
		return 0, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return entity.UserId, nil
}

func (s *Service) LoginUser(ctx context.Context, params dto.AuthLoginRequest) (bool, error) {
	// Fetch the user from the repository based on the username
	user, err := s.repo.GetUserByUsername(ctx, params.Username)
	if err != nil {
		return false, err
	}

	// Check if the user exists
	if user == nil {
		return false, nil // User not found
	}

	// Compare the provided password with the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		// Passwords don't match
		return false, nil
	}

	// Passwords match, user is authenticated
	return true, nil
}
