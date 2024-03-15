package product

import (
	"context"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) IndexProduct(ctx context.Context) ([]model.Product, error) {
	return []model.Product{}, ctx.Err()
}

func (s Service) StoreProduct(ctx context.Context, params *model.Product) (string, error) {
	return "", ctx.Err()
}

func (s Service) ShowProduct(ctx context.Context, id int) (model.Product, error) {
	return model.Product{}, ctx.Err()
}

func (s Service) UpdateProduct(ctx context.Context, params *model.Product) (string, error) {
	return "", ctx.Err()
}

func (s Service) DeleteProduct(ctx context.Context, id int) (string, error) {
	return "", ctx.Err()
}
