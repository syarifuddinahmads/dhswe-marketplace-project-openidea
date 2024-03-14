package todo

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/db"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

type CreateParams struct {
	Name        string       `valid:"required"`
	Description string       `valid:"required"`
	Status      model.Status `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (int, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, utils.ErrArgument{Wrapped: err}
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.ToDo{
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
		CreatedOn:   time.Now().UTC(),
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return entity.ID, err
}

func (s Service) Delete(ctx context.Context, id int) error {
	todo, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	now := time.Now().UTC()
	todo.DeletedOn = &now
	err = s.repo.Update(ctx, todo)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (s Service) Get(ctx context.Context, id int) (model.ToDo, error) {
	todo, err := s.repo.Find(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.ToDo{}, utils.ErrArgument{Wrapped: errors.New("todo object not found")}
	default:
		return model.ToDo{}, err
	}
	return todo, nil
}

func (s Service) GetAll(ctx context.Context) ([]model.ToDo, error) {
	todos, err := s.repo.FindAll(ctx)
	switch {
	case err == nil:
		return todos, nil
	case errors.Is(err, err):
		return nil, utils.ErrArgument{Wrapped: errors.New("todo object not found")}
	default:
		return nil, err
	}
}

type UpdateParams struct {
	ID          int `valid:"required"`
	Name        *string
	Description *string
	Status      *model.Status
}

func (s Service) Update(ctx context.Context, params UpdateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return utils.ErrArgument{Wrapped: err}
	}

	// find todo object
	todo, err := s.Get(ctx, params.ID)
	if err != nil {
		return err
	}

	if params.Name != nil {
		todo.Name = *params.Name
	}
	if params.Description != nil {
		todo.Description = *params.Description
	}
	if params.Status != nil {
		if !params.Status.IsValid() {
			return utils.ErrArgument{Wrapped: errors.New("given status not valid")}
		}
		todo.Status = *params.Status
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = s.repo.Update(ctx, todo)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
