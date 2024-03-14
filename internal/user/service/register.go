package service

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/user/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

type CreateParams struct {
	Name     string `valid:"required"`
	Username string `valid:"required"`
	Password string `valid:"required"`
}

func (s Service) Register(ctx context.Context, params CreateParams) (string, string, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return "", "", utils.ErrArgument{Wrapped: err}
	}
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return "", "", err
	}
	defer tx.Rollback()
	entity := model.User{
		Name:     params.Name,
		Username: params.Username,
		Password: params.Password,
	}

	err = s.repo.Register(ctx, &entity)
	if err != nil {
		return "", "", err
	}
	token, err := entity.GenerateToken()
	if err != nil {
		return "", "", err
	}
	err = tx.Commit()
	return entity.UserId, token, err
}
