package service

import (
	"context"
	"errors"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/todo/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/db"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

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
