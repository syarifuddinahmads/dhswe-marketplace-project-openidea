package product

import (
	"context"
	"fmt"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/response"
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

func (s Service) StoreProduct(ctx context.Context, payload *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	_, err := s.repo.CreateProduct(ctx, payload)
	fmt.Print("Service Create Product : ")
	fmt.Print(err)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result := &dto.CreateProductResponse{
		Response: "product added successfully",
	}

	return result, nil
}

func (s Service) ShowProduct(ctx context.Context, id int) (model.Product, error) {
	return model.Product{}, ctx.Err()
}

func (s Service) UpdateProduct(ctx context.Context, productID string, payload *dto.UpdateProductRequest) error {
	// Call the repository method to update the product
	err := s.repo.UpdateProduct(ctx, productID, payload)
	fmt.Print("Update Product : ")
	fmt.Println(err)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	return nil
}

func (s Service) DeleteProduct(ctx context.Context, productID string) error {
	err := s.repo.DeleteProduct(ctx, productID)
	fmt.Print("Delete Product : ")
	fmt.Println(err)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	return nil
}
