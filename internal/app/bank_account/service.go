package bank_account

import (
	"context"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/response"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s *Service) CreateBank(ctx context.Context, payload *dto.BankAccountRequest) (*dto.BankAccountResponse, error) {
	// Call the repository Register function
	data, err := s.repo.CreateBank(ctx, payload)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Service bank 1")
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result := &dto.BankAccountResponse{
		Bank_Name: data.BankName,
	}

	return result, nil
}

// func (s Service) Delete(ctx context.Context, id int) error {
// 	todo, err := s.Get(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := s.repo.Db.BeginTxx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	// Defer a rollback in case anything fails.
// 	defer tx.Rollback()

// 	now := time.Now().UTC()
// 	todo.DeletedOn = &now
// 	err = s.repo.Update(ctx, todo)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	return err
// }

// func (s Service) Get(ctx context.Context, id int) (model.ToDo, error) {
// 	todo, err := s.repo.Find(ctx, id)
// 	switch {
// 	case err == nil:
// 	case errors.As(err, nil):
// 		return model.ToDo{}, utils.ErrArgument{Wrapped: errors.New("todo object not found")}
// 	default:
// 		return model.ToDo{}, err
// 	}
// 	return todo, nil
// }

// func (s Service) GetAll(ctx context.Context) ([]model.ToDo, error) {
// 	todos, err := s.repo.FindAll(ctx)
// 	switch {
// 	case err == nil:
// 		return todos, nil
// 	case errors.Is(err, err):
// 		return nil, utils.ErrArgument{Wrapped: errors.New("todo object not found")}
// 	default:
// 		return nil, err
// 	}
// }

// type UpdateParams struct {
// 	ID          int `valid:"required"`
// 	Name        *string
// 	Description *string
// 	Status      *model.Status
// }

// func (s Service) Update(ctx context.Context, params UpdateParams) error {
// 	if _, err := govalidator.ValidateStruct(params); err != nil {
// 		return utils.ErrArgument{Wrapped: err}
// 	}

// 	// find todo object
// 	todo, err := s.Get(ctx, params.ID)
// 	if err != nil {
// 		return err
// 	}

// 	if params.Name != nil {
// 		todo.Name = *params.Name
// 	}
// 	if params.Description != nil {
// 		todo.Description = *params.Description
// 	}
// 	if params.Status != nil {
// 		if !params.Status.IsValid() {
// 			return utils.ErrArgument{Wrapped: errors.New("given status not valid")}
// 		}
// 		todo.Status = *params.Status
// 	}

// 	tx, err := s.repo.Db.BeginTxx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	// Defer a rollback in case anything fails.
// 	defer tx.Rollback()

// 	err = s.repo.Update(ctx, todo)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	return err
// }
