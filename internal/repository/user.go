package repository

import (
	"context"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/db"
)

func (r Repository) Register(ctx context.Context, entity *model.User) error {
	query := `INSERT INTO users (name, username, password)
				VALUES (:name, :username, :password);`
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

func (r Repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `SELECT * FROM users WHERE username = ?`

	var user model.User
	err := r.Db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, db.HandleError(err)
	}

	return &user, nil
}
