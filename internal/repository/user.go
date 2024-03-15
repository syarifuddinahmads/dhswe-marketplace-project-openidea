package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

func (r Repository) Register(ctx context.Context, entity *dto.AuthRegisterRequest) (*model.User, error) {
	query := `INSERT INTO users (name, username, password)
				VALUES (:name, :username, :password);`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)

	var user model.User
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	// Prepare query
	query := `SELECT user_id, name, username, password FROM users WHERE username = $1`

	// Execute the query
	row := r.Db.QueryRowContext(ctx, query, username)

	// Scan the result
	var user model.User
	err := row.Scan(&user.UserId, &user.Name, &user.Username, &user.Password)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("User not found")
			return nil, fmt.Errorf("user with username %s not found", username)
		}
		return nil, err
	}

	return &user, nil
}
