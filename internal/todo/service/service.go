package service

import (
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/todo/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}
