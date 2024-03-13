package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

type User interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.User, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.User, error)
	FindByEmail(ctx context.Context, email *string) (*model.User, error)
}

type user struct {
	Db *sql.DB
}

func NewUser(db *sql.DB) *user {
	return &user{
		db,
	}
}

func (r *user) FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.User, *dto.PaginationInfo, error) {
	var users []model.User
	var count int64

	query := "SELECT * FROM users"

	if payload.Search != "" {
		//search := "%" + strings.ToLower(payload.Search) + "%"
		query += " WHERE lower(name) LIKE $1 OR lower(email) LIKE $1"
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS count", query)
	err := r.Db.QueryRowContext(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(p)
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Username) // Assuming the columns are ID, Name, Email
		if err != nil {
			return nil, nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return users, dto.CheckInfoPagination(p, count), nil
}

func (r *user) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE id = $1"
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Username) // Assuming the columns are ID, Name, Email
	if err == sql.ErrNoRows {
		return model.User{}, fmt.Errorf("user not found")
	}
	return user, err
}

func (r *user) FindByEmail(ctx context.Context, email *string) (*model.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	var user model.User
	err := r.Db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Username) // Assuming the columns are ID, Name, Email
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
