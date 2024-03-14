package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/bank/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/db"
)

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Create(ctx context.Context, entity *model.BankAcccount) error {
	query := `INSERT INTO todo (bank_name, bank_account_name, bank_account_number,)
				VALUES (:bank_name, :bank_account_name, :bank_account_number) RETURNING bank_account_id;`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return db.HandleError(err)
		}
	}
	return db.HandleError(err)
}
