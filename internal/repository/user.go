package repository

import (
	"context"
	"fmt"
	// "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	// "strings"

	"github.com/jmoiron/sqlx"
	// "gorm.io/gorm"
)

type User interface {
	// FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.User, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.User, error)
	// FindByEmail(ctx context.Context, email *string) (*model.User, error)
	FindByUsername(ctx context.Context, username *string) (*model.User, error)
}

// type user struct {
// 	Db *gorm.DB
// }
// func NewUser(db *gorm.DB) *user {
// 	return &user{
// 		db,
// 	}
// }
type user struct {
	Db *sqlx.DB
}

func NewUser(db *sqlx.DB) *user {
	return &user{
		db,
	}
}

// func (r *user) FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.User, *dto.PaginationInfo, error) {
// 	var users []model.User
// 	var count int64

// 	query := r.Db.WithContext(ctx).Model(&model.User{})

// 	if payload.Search != "" {
// 		search := "%" + strings.ToLower(payload.Search) + "%"
// 		query = query.Where("lower(name) LIKE ? or lower(email) Like ? ", search, search)
// 	}

// 	countQuery := query
// 	if err := countQuery.Count(&count).Error; err != nil {
// 		return nil, nil, err
// 	}

// 	limit, offset := dto.GetLimitOffset(p)

// 	err := query.Limit(limit).Offset(offset).Find(&users).Error

// 	return users, dto.CheckInfoPagination(p, count), err
// }

func (r *user) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	err := r.Db.GetContext(ctx, &user, fmt.Sprintf("SELECT user_id, username, name from users where user_id = %d", id), id)

	// err := r.Db.WithContext(ctx).Model(&model.User{}).QueryRow("SELECT user_id, username, name from users where user_id = ?").Scan(&user).Error
	return user, err
}

// func (r *user) FindByEmail(ctx context.Context, email *string) (*model.User, error) {
// 	conn := r.Db.WithContext(ctx)

// 	var data model.User
// 	err := conn.Where("email = ?", email).First(&data).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

func (r *user) FindByUsername(ctx context.Context, username *string) (*model.User, error) {
	// conn := r.Db.WithContext(ctx)
	var user model.User
	err := r.Db.GetContext(ctx, &user, fmt.Sprintf("SELECT user_id, username, name from users where username = %d", username), username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
