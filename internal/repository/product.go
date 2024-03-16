package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

func (r Repository) FindProduct(ctx context.Context, id int) (model.Product, error) {
	entity := model.Product{}
	query := fmt.Sprintf(
		"SELECT * FROM product WHERE id = $1 AND deleted_at IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, err
}

// func (r Repository) CreateProduct(ctx context.Context, entity *dto.CreateProductRequest) (*model.Product, error) {
// 	query := `INSERT INTO product (name, price, image_url, condition, tags, is_purchaseable, created_at, updated_at)
// 				VALUES (:name, :price, :image_url, :condition, :tags, :is_purchaseable, :created_at, :updated_at) RETURNING id;`
// 	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
// 	var product model.Product
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		err = rows.StructScan(entity)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return &product, err
// }

func (r Repository) CreateProduct(ctx context.Context, entity *dto.CreateProductRequest) (*model.Product, error) {

	tagsArrayLiteral := "{" + strings.Join(entity.Tags, ",") + "}"

	query := `INSERT INTO products (name, price, image_url,stock, condition, tags, is_purchaseable, created_at, updated_at)
				VALUES (:name, :price, :image_url,:stock, :condition, :tags, :is_purchaseable, :created_at, :updated_at);`

	rows, err := r.Db.NamedQueryContext(ctx, query, map[string]interface{}{
		"name":            entity.Name,
		"price":           entity.Price,
		"image_url":       entity.Image_Url,
		"stock":           entity.Stock,
		"condition":       entity.Condition,
		"tags":            tagsArrayLiteral, // Use the comma-separated string here
		"is_purchaseable": entity.Is_Purchaseable,
		"created_at":      time.Now(),
		"updated_at":      time.Now(),
	})

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product model.Product
	for rows.Next() {
		err = rows.Scan(&product.Product_Id)
		if err != nil {
			return nil, err
		}
	}

	return &product, nil
}

func (r Repository) UpdateProduct(ctx context.Context, productID string, payload *dto.UpdateProductRequest) error {
	tagsArrayLiteral := "{" + strings.Join(payload.Tags, ",") + "}"

	// Construct the update query
	query := `UPDATE products 
			  SET name=:name, price=:price, image_url=:image_url, stock=:stock, condition=:condition, tags=:tags, is_purchaseable=:is_purchaseable, updated_at=:updated_at 
			  WHERE product_id=:product_id;`

	// Execute the update query
	_, err := r.Db.NamedExecContext(ctx, query, map[string]interface{}{
		"product_id":      productID,
		"name":            payload.Name,
		"price":           payload.Price,
		"image_url":       payload.Image_Url,
		"stock":           payload.Stock,
		"condition":       payload.Condition,
		"tags":            tagsArrayLiteral,
		"is_purchaseable": payload.Is_Purchaseable,
		"updated_at":      time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetAllProduct(ctx context.Context) ([]model.Product, error) {
	var entities []model.Product
	query := fmt.Sprintf(
		"SELECT * FROM products WHERE deleted_at IS NULL",
	)
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, err
}

func (r Repository) DeleteProduct(ctx context.Context, productID string) error {
	query := fmt.Sprintf(
		"DELETE FROM products WHERE product_id = $1",
	)
	fmt.Println("Executing SQL query:", query) // Add this line to log the SQL query
	_, err := r.Db.ExecContext(ctx, query, productID)
	return err
}
