package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

type Product interface {
	Create(ctx context.Context, data model.Product) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Product, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Product, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type product struct {
	Db *sql.DB
}

func NewProduct(db *sql.DB) *product {
	return &product{
		db,
	}
}

func (p *product) Create(ctx context.Context, data model.Product) error {
	query := `INSERT INTO products (name, stock, description) VALUES ($1, $2, $3)`
	_, err := p.Db.ExecContext(ctx, query, data.Name, data.Stock, data.Description)
	return err
}

func (p *product) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Product, *dto.PaginationInfo, error) {
	var products []model.Product
	var count int64

	query := `SELECT * FROM products`

	if payload.Search != "" {
		// search := "%" + strings.ToLower(payload.Search) + "%"
		query += " WHERE LOWER(name) LIKE $1"
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS count", query)
	err := p.Db.QueryRowContext(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := p.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Stock, &product.Description)
		if err != nil {
			return nil, nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return products, dto.CheckInfoPagination(paginate, count), nil
}

func (p *product) FindByID(ctx context.Context, ID uint) (model.Product, error) {
	var data model.Product
	query := `SELECT * FROM products WHERE id = $1`
	err := p.Db.QueryRowContext(ctx, query, ID).Scan(&data.ID, &data.Name, &data.Stock, &data.Description)
	if err == sql.ErrNoRows {
		return model.Product{}, fmt.Errorf("product not found")
	}
	return data, err
}

func (p *product) Update(ctx context.Context, ID uint, data map[string]interface{}) error {
	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data)+1)

	for k, v := range data {
		columns = append(columns, fmt.Sprintf("%s=$%d", k, len(values)+1))
		values = append(values, v)
	}

	values = append(values, ID)
	query := fmt.Sprintf("UPDATE products SET %s WHERE id=$%d", strings.Join(columns, ", "), len(values))

	_, err := p.Db.ExecContext(ctx, query, values...)
	return err
}

func (p *product) Delete(ctx context.Context, ID uint) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := p.Db.ExecContext(ctx, query, ID)
	return err
}
