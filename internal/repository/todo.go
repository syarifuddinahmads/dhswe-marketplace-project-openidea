package repository

import (
	"context"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

func (r Repository) Find(ctx context.Context, id int) (model.ToDo, error) {
	entity := model.ToDo{}
	query := fmt.Sprintf(
		"SELECT * FROM todo WHERE id = $1 AND deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, err
}

func (r Repository) Create(ctx context.Context, entity *model.ToDo) error {
	query := `INSERT INTO todo (name, description, status, created_on, updated_on)
				VALUES (:name, :description, :status, :created_on, :updated_on) RETURNING id;`
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

func (r Repository) Update(ctx context.Context, entity model.ToDo) error {
	query := `UPDATE todo
    		  	SET name = :name, 
    		  	    description = :description, 
    		  	    status = :status, 
    		  	    created_on = :created_on, 
    		  	    updated_on = :updated_on, 
    		  	    deleted_on = :deleted_on
				WHERE id = :id;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return err
}

func (r Repository) FindAll(ctx context.Context) ([]model.ToDo, error) {
	var entities []model.ToDo
	query := fmt.Sprintf(
		"SELECT * FROM todo WHERE deleted_on IS NULL",
	)
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, err
}
