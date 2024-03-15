package repository

import (
	"context"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

func (r Repository) FindProduct(ctx context.Context, id int) (model.Product, error) {
	entity := model.Product{}
	query := fmt.Sprintf(
		"SELECT * FROM product WHERE id = $1 AND deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, err
}

func (r Repository) CreateProduct(ctx context.Context, entity *model.Product) error {
	query := `INSERT INTO product (name, price, imageUrl, condition, tags, isPurchasable, created_on, updated_on)
				VALUES (:name, :price, :imageUrl, :condition, :tags, :isPurchasable, :created_on, :updated_on) RETURNING id;`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return err
		}
	}
	return err
}

func (r Repository) UpdateProduct(ctx context.Context, entity model.Product) error {
	query := `UPDATE product
    		  	SET name = :name, 
    		  	    price = :price, 
    		  	    imageUrl = :imageUrl, 
    		  	    condition = :condition, 
    		  	    tags = :tags, 
    		  	    isPurchasable = :isPurchasable, 
    		  	    updated_on = :updated_on, 
    		  	    deleted_on = :deleted_on
				WHERE id = :id;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return err
}

func (r Repository) FindAllProduct(ctx context.Context) ([]model.Product, error) {
	var entities []model.Product
	query := fmt.Sprintf(
		"SELECT * FROM product WHERE deleted_on IS NULL",
	)
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, err
}

func (r Repository) DeleteProduct(ctx context.Context, id int) error {
	query := fmt.Sprintf("UPDATE product SET deleted_on = NOW() WHERE id = $1")
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}
