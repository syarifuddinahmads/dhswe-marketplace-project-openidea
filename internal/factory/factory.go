package factory

import (
	database "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/db"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
)

type Factory struct {
	ProductRepository repository.Product
	UserRepository    repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		ProductRepository: repository.NewProduct(db),
		UserRepository:    repository.NewUser(db),
	}
}
