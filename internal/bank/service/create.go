package service

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/bank/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

type CreateParams struct {
	BankName          string `valid:"required"`
	BankAccountName   string `valid:"required"`
	BankAccountNumber string `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (string, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return "", utils.ErrArgument{Wrapped: err}
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return "", err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.BankAcccount{
		BankName:          params.BankName,
		BankAccountName:   params.BankAccountName,
		BankAccountNumber: params.BankAccountNumber,
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	return entity.BankAccountName, err
}
