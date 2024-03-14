package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/user/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/db"
)

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Register(ctx context.Context, entity *model.User) error {
	query := `INSERT INTO user (name, email, password)
				VALUES (:name, :email, :password);`
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
