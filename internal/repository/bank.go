package repository

import (
	"context"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
)

// func (r Repository) Find(ctx context.Context, id int) (model.ToDo, error) {
// 	entity := model.ToDo{}
// 	query := fmt.Sprintf(
// 		"SELECT * FROM todo WHERE id = $1 AND deleted_on IS NULL",
// 	)
// 	err := r.Db.GetContext(ctx, &entity, query, id)
// 	return entity, err
// }

func (r Repository) CreateBank(ctx context.Context, entity *dto.BankAccountRequest) (*model.BankAccount, error) {
	query := `INSERT INTO bank_accounts (bank_name, bank_account_name, bank_account_number)
				VALUES (:bank_name, :bank_account_name, :bank_account_number);`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	var bank model.BankAccount
	if err != nil {
		fmt.Println("Repo bank 1")
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			fmt.Println("Repo bank 2")

			return nil, err
		}
	}
	return &bank, nil
}

func (r Repository) UpdateBank(ctx context.Context, entity model.BankAccount) error {
	query := `UPDATE bank_accounts
    		  	SET bank_name = :bank_name, 
    		  	    bank_account_name = :bank_account_name, 
    		  	    bank_account_number = :bank_account_number,
				WHERE bank_account_id = :bank_account_id;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return err
}

// func (r Repository) FindAll(ctx context.Context) ([]model.ToDo, error) {
// 	var entities []model.ToDo
// 	query := fmt.Sprintf(
// 		"SELECT * FROM todo WHERE deleted_on IS NULL",
// 	)
// 	err := r.Db.SelectContext(ctx, &entities, query)
// 	return entities, err
// }
